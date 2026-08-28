package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/assessment"
	"quantumwizard.hu/qwsg/internal/installer"
)

func TestFreshGuardianEvidenceWaitImmediateDelayedAndPreservedStale(t *testing.T) {
	tests := []struct {
		name     string
		ids      []string
		previous string
		pauses   int
	}{
		{name: "clean immediate", ids: []string{"new-clean"}},
		{name: "clean delayed", ids: []string{"", "", "new-clean"}, pauses: 2},
		{name: "preserved-state reinstall", previous: "old-install", ids: []string{"old-install", "old-install", "new-install"}, pauses: 2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			index, pauses := 0, 0
			probe := func() (string, bool) {
				id := tc.ids[index]
				if index < len(tc.ids)-1 {
					index++
				}
				return id, id != ""
			}
			pause := func(context.Context, time.Duration) error { pauses++; return nil }
			if err := waitForFreshGuardianEvidence(context.Background(), tc.previous, time.Second, probe, pause); err != nil || pauses != tc.pauses {
				t.Fatalf("err=%v pauses=%d", err, pauses)
			}
		})
	}
}

func TestFreshGuardianEvidenceWaitTimeoutAndCancellationAreBounded(t *testing.T) {
	probeCalls := 0
	probe := func() (string, bool) { probeCalls++; return "preserved", true }
	pause := func(ctx context.Context, _ time.Duration) error { <-ctx.Done(); return ctx.Err() }
	if err := waitForFreshGuardianEvidence(context.Background(), "preserved", 2*time.Millisecond, probe, pause); !errors.Is(err, errGuardianEvidenceTimeout) || probeCalls != 1 {
		t.Fatalf("timeout err=%v calls=%d", err, probeCalls)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := waitForFreshGuardianEvidence(ctx, "preserved", time.Second, probe, pause); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation err=%v", err)
	}
}

func TestGuidedCompletionOptionalPartialRenders97And100InAllLocales(t *testing.T) {
	for _, language := range []installer.Language{installer.English, installer.Hungarian, installer.German} {
		progress, render, complete, rendered := guidedCompletionFixture()
		var out, errout bytes.Buffer
		report := assessment.Report{Domains: []assessment.DomainSummary{{Domain: "guardian_service", State: assessment.Ready}, {Domain: "notification", State: assessment.Partial}, {Domain: "overall", State: assessment.Partial}}}
		code := finishGuidedInstallation(installer.Catalog{Language: language}, &progress, true, "notify", report, &out, &errout, render, complete)
		text := out.String()
		if code != 0 || errout.Len() != 0 || !strings.Contains(rendered.String(), "completion:97") || !strings.Contains(text, "100%") || !strings.Contains(text, version) || !strings.Contains(text, installer.Catalog{Language: language}.Text("summary.partial")) {
			t.Fatalf("language=%s code=%d rendered=%q out=%q err=%q", language, code, rendered.String(), text, errout.String())
		}
	}
}

func TestGuidedCompletionMandatoryNotReadyStaysExit4WithoutFalseSummary(t *testing.T) {
	progress, render, complete, rendered := guidedCompletionFixture()
	var out, errout bytes.Buffer
	report := assessment.Report{Domains: []assessment.DomainSummary{{Domain: "guardian_service", State: assessment.NotReady}, {Domain: "overall", State: assessment.NotReady}}}
	code := finishGuidedInstallation(installer.Catalog{Language: installer.English}, &progress, true, "notify", report, &out, &errout, render, complete)
	if code != 4 || !strings.Contains(errout.String(), "Readiness: incomplete") || strings.Contains(out.String(), "100%") || strings.Contains(rendered.String(), "completion") {
		t.Fatalf("code=%d rendered=%q out=%q err=%q", code, rendered.String(), out.String(), errout.String())
	}
}

func TestReadinessTimeoutDiagnosticsAreLocalizedAndSecretFree(t *testing.T) {
	for _, language := range []installer.Language{installer.English, installer.Hungarian, installer.German} {
		message := installer.Catalog{Language: language}.Text("readiness.timeout")
		if !strings.Contains(message, "qwsg readiness") || !strings.Contains(message, "qwsg setup") {
			t.Fatalf("language=%s message=%q", language, message)
		}
		for _, secret := range []string{"password=", "token=", "credential=", "private key"} {
			if strings.Contains(strings.ToLower(message), secret) {
				t.Fatalf("language=%s leaked %q", language, secret)
			}
		}
	}
}

func guidedCompletionFixture() (installer.Progress, func(installer.PhaseID), func(installer.PhaseID), *bytes.Buffer) {
	progress := installer.NewProgress()
	for _, phase := range installer.Phases[:7] {
		_ = progress.Start(phase.ID)
		_ = progress.Complete(phase.ID)
	}
	_ = progress.Start(installer.PhaseReadiness)
	rendered := &bytes.Buffer{}
	render := func(phase installer.PhaseID) {
		_ = progress.Start(phase)
		fmt.Fprintf(rendered, "%s:%d\n", phase, progress.Percent())
	}
	complete := func(phase installer.PhaseID) { _ = progress.Complete(phase) }
	return progress, render, complete, rendered
}
