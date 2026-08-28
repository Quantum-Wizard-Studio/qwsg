package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/changenotification"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/credentialstore"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/smtpnotification"
)

var changeDispatcher changenotification.Dispatcher
var managedChangeDelivery = deliverManagedChange

type smtpChangeSender struct{ provider smtpnotification.Provider }

func (s smtpChangeSender) Send(subject, body string) error {
	result := s.provider.DeliverText(context.Background(), subject, body)
	if result.Status != notification.StatusAccepted && result.Status != notification.StatusDelivered {
		return fmt.Errorf("delivery failed")
	}
	return nil
}

func deliverManagedChange(event changenotification.Event, out io.Writer) changenotification.DeliveryResult {
	path, err := configurationstore.SelectPath("", os.Getenv)
	if err != nil {
		return reportChangeDelivery(out, changenotification.DeliveryDisabled)
	}
	source, found, err := configurationstore.Load(path)
	if err != nil || !found {
		return reportChangeDelivery(out, changenotification.DeliveryDisabled)
	}
	effective, err := resolveLocalConfiguration(source, true, nil)
	if err != nil {
		return reportChangeDelivery(out, changenotification.DeliveryFailed)
	}
	cfg, err := smtpnotification.FromEffective(effective)
	if err != nil || !cfg.Enabled || !cfg.LifecycleEnabled {
		return reportChangeDelivery(out, changenotification.DeliveryDisabled)
	}
	password := []byte(nil)
	if cfg.Auth == "password" {
		password, err = credentialstore.Load(path, cfg.CredentialRef)
		if err != nil {
			return reportChangeDelivery(out, changenotification.DeliveryFailed)
		}
	}
	return reportChangeDelivery(out, changeDispatcher.Deliver(true, effective.Values.Locale, event, smtpChangeSender{smtpnotification.Provider{Config: cfg, Password: password}}))
}

func reportChangeDelivery(out io.Writer, result changenotification.DeliveryResult) changenotification.DeliveryResult {
	fmt.Fprintf(out, "Admin notification: %s\n", result)
	return result
}

func managedEvent(kind changenotification.EventType, result changenotification.OperationResult, id, previous, next, reason string) changenotification.Event {
	host, _ := os.Hostname()
	host = strings.TrimSpace(host)
	if host == "" || strings.ContainsAny(host, "\r\n") {
		host = "local-host"
	}
	return changenotification.Event{ID: id, Host: host, Type: kind, Result: result, PreviousVersion: previous, NewVersion: next, Reason: reason, At: time.Now().UTC(), ActionRequired: result == changenotification.Failed}
}
