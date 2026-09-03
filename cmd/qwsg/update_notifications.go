package main

import (
	"context"

	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/smtpnotification"
)

type smtpUpdateSender struct{ provider smtpnotification.Provider }

func (s smtpUpdateSender) DeliverText(ctx context.Context, subject, body string) bool {
	result := s.provider.DeliverText(ctx, subject, body)
	return result.Status == notification.StatusAccepted || result.Status == notification.StatusDelivered
}

func updateNotificationEnabled(effective configuration.Effective) bool {
	for _, extension := range effective.Values.Extensions {
		if extension.ID == "installer.update-policy" {
			return extension.Fields["policy"] == "notify"
		}
	}
	return false
}
