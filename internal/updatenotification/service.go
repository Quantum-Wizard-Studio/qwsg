// Package updatenotification turns an authenticated actionable awareness
// result into one operator notification. It never discovers or installs.
package updatenotification

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"quantumwizard.hu/qwsg/internal/updateawareness"
)

const ProductionSource = "https://releases.quantumwizard.hu/qwsg/v1/release-index.json"

type Store interface {
	RecordNotification(string, time.Time) error
}

type Sender interface {
	DeliverText(context.Context, string, string) bool
}

type Service struct {
	Enabled   bool
	Locale    string
	Source    string
	Store     Store
	Sender    Sender
	Now       func() time.Time
	mu        sync.Mutex
	delivered map[string]bool
}

func (s *Service) Notify(ctx context.Context, state updateawareness.State) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	identity, eligible := Eligible(state)
	if !s.Enabled || !eligible {
		return nil
	}
	if state.LastNotification != nil && state.LastNotification.Identity == identity {
		return nil
	}
	if s.delivered[identity] {
		return nil
	}
	if s.Store == nil || s.Sender == nil {
		return fmt.Errorf("update notification unavailable")
	}
	subject, body := Render(s.Locale, state, s.Source)
	if !s.Sender.DeliverText(ctx, subject, body) {
		return fmt.Errorf("update notification delivery failed")
	}
	now := time.Now().UTC()
	if s.Now != nil {
		now = s.Now().UTC()
	}
	if err := s.Store.RecordNotification(identity, now); err != nil {
		return err
	}
	if s.delivered == nil {
		s.delivered = make(map[string]bool)
	}
	s.delivered[identity] = true
	return nil
}

func Eligible(state updateawareness.State) (string, bool) {
	if updateawareness.Validate(state) != nil || state.Status != updateawareness.UpdateAvailable || state.LastAttempt.Outcome != updateawareness.AttemptSuccess || state.LastSuccess == nil {
		return "", false
	}
	o := state.LastSuccess
	if o.Channel != "stable" || !o.TransportAuthenticated || o.Authenticity.Scheme != "ed25519" || o.ReleaseStatus != "active" {
		return "", false
	}
	return updateawareness.NotificationIdentity(state.SourceID, state.Channel, o.ReleaseVersion, o.ArtifactSHA256, o.Authenticity.KeyID), true
}

func Render(locale string, state updateawareness.State, source string) (string, string) {
	if source == "" {
		source = ProductionSource
	}
	lang := strings.ToLower(strings.SplitN(locale, "-", 2)[0])
	o := state.LastSuccess
	if lang == "hu" {
		return "[QWSG] Hitelesített frissítés érhető el", fmt.Sprintf("QWSG-frissítés érhető el.\nTelepített verzió: %s\nElérhető verzió: %s\nCsatorna: stable\nHitelesség: a kiadási metaadat Ed25519-aláírása ellenőrizve\nForrás: %s\nA telepítés nem automatikus. Következő lépés: ellenőrizze a `qwsg update status` kimenetét, majd indítsa a frissítést saját döntése szerint.\n", state.Installed.Version, o.ReleaseVersion, source)
	}
	return "[QWSG] Authenticated update available", fmt.Sprintf("A QWSG update is available.\nInstalled version: %s\nAvailable version: %s\nChannel: stable\nAuthenticity: release metadata Ed25519 signature verified\nSource: %s\nInstallation is not automatic. Next step: review `qwsg update status`, then start the update when you choose.\n", state.Installed.Version, o.ReleaseVersion, source)
}
