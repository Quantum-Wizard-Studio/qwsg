package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/credentialstore"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/smtpnotification"
)

func runNotification(args []string, out, errout io.Writer) int {
	if len(args) == 0 || isHelp(args[0]) {
		fmt.Fprintln(out, "Usage: qwsg notification <preflight|test|credential set --from-file FILE> [--config FILE] [--format human|json]")
		return 0
	}
	action := args[0]
	rest := args[1:]
	if action == "credential" {
		return runCredential(rest, out, errout)
	}
	options, err := parseConfigOptions(rest)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	path, err := configurationstore.SelectPath(options.path, os.Getenv)
	if err != nil {
		return configFailure(errout, err)
	}
	source, found, err := configurationstore.Load(path)
	if err != nil || !found {
		if err == nil {
			err = configurationstore.ErrUnavailable
		}
		return configFailure(errout, err)
	}
	effective, err := resolveLocalConfiguration(source, true, nil)
	if err != nil {
		return configFailure(errout, err)
	}
	cfg, err := smtpnotification.FromEffective(effective)
	if err != nil {
		return configFailure(errout, configurationstore.ErrInvalid)
	}
	password := []byte(nil)
	available := cfg.Auth != "password"
	if cfg.Auth == "password" {
		password, err = credentialstore.Load(path, cfg.CredentialRef)
		available = err == nil
	}
	findings := smtpnotification.Preflight(cfg, available)
	if action == "preflight" {
		if options.format == formatJSON {
			return writeJSON(out, errout, map[string]any{"status": "assessed", "findings": findings})
		}
		for _, f := range findings {
			fmt.Fprintf(out, "%s: %s\n", safeText(f.Requirement), f.Classification)
		}
		if !smtpnotification.Ready(findings) {
			return 1
		}
		return 0
	}
	if action != "test" {
		return usageError(errout, "unknown notification operation")
	}
	if !cfg.Enabled || !smtpnotification.Ready(findings) {
		fmt.Fprintln(errout, "notification test failed: notification_not_ready")
		return 1
	}
	now := time.Now().UTC()
	req := notification.Request{DestinationRef: cfg.Recipient, Deadline: now.Add(cfg.Timeout), Envelope: notification.DeliveryEnvelope{AlertRecordID: "test", LifecycleID: "test", ConditionKey: "test", Event: alert.EventEntered, Severity: alert.Warning, Category: alert.EngineeringCondition, ReasonToken: "operator_requested_test", EventTime: now, EvidenceReferences: []string{}}}
	result := smtpnotification.Provider{Config: cfg, Password: password}.Deliver(context.Background(), req)
	if result.Status != notification.StatusAccepted && result.Status != notification.StatusDelivered {
		fmt.Fprintln(errout, "notification test failed: smtp_delivery_failed")
		return 1
	}
	fmt.Fprintln(out, "QWSG test notification accepted by the configured SMTP server.")
	return 0
}

func runCredential(args []string, out, errout io.Writer) int {
	if len(args) < 3 || args[0] != "set" || args[1] != "--from-file" {
		return usageError(errout, "credential update requires: credential set --from-file FILE")
	}
	input := args[2]
	options, err := parseConfigOptions(args[3:])
	if err != nil {
		return usageError(errout, "%v", err)
	}
	info, err := os.Lstat(input)
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm() != 0600 {
		fmt.Fprintln(errout, "credential update failed: input_file_unsafe")
		return 1
	}
	secret, err := os.ReadFile(input)
	if err != nil {
		return 1
	}
	path, err := configurationstore.SelectPath(options.path, os.Getenv)
	if err != nil {
		return configFailure(errout, err)
	}
	source, found, err := configurationstore.Load(path)
	if err != nil || !found {
		return configFailure(errout, configurationstore.ErrUnavailable)
	}
	effective, err := resolveLocalConfiguration(source, true, nil)
	if err != nil {
		return configFailure(errout, err)
	}
	cfg, err := smtpnotification.FromEffective(effective)
	if err != nil || cfg.CredentialRef == "" {
		return configFailure(errout, configurationstore.ErrInvalid)
	}
	if err = credentialstore.Save(path, cfg.CredentialRef, secret); err != nil {
		if errors.Is(err, credentialstore.ErrPermission) || errors.Is(err, credentialstore.ErrUnsafe) {
			fmt.Fprintln(errout, "credential update failed: credential_path_unsafe")
		} else {
			fmt.Fprintln(errout, "credential update failed: credential_unavailable")
		}
		return 1
	}
	fmt.Fprintln(out, "SMTP credential stored safely.")
	return 0
}
