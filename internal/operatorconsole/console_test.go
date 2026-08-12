package operatorconsole

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

func unavailable(t *testing.T) presentationmodel.Overview {
	t.Helper()
	now := time.Date(2026, 8, 7, 20, 0, 0, 0, time.UTC)
	value, err := presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: now, FreshForNS: int64(time.Hour)})
	if err != nil {
		t.Fatal(err)
	}
	return value
}

func TestHomeIsOverviewOnlyLocalizedAndDeterministic(t *testing.T) {
	for _, locale := range []string{"en", "hu"} {
		state, err := NewState(unavailable(t), locale, Capabilities{Width: 80, Height: 30})
		if err != nil {
			t.Fatal(err)
		}
		first, second := Render(state), Render(state)
		if first != second {
			t.Fatal("render is not deterministic")
		}
		for _, value := range []string{string(state.Overview.Condition), string(state.Overview.Attention), string(state.Overview.Guardian)} {
			if !strings.Contains(first, text(locale, value)) {
				t.Fatalf("%s missing %s", locale, value)
			}
		}
		if strings.Contains(first, state.Overview.ID) {
			t.Fatal("home exposed canonical ID")
		}
	}
}

func TestNavigationBoundariesAndEveryScreen(t *testing.T) {
	state, _ := NewState(unavailable(t), "en", Capabilities{Width: 80, Height: 30})
	state = Transition(state, Up)
	if state.Selection != 0 {
		t.Fatal("up underflow")
	}
	for range 10 {
		state = Transition(state, Down)
	}
	if state.Selection != 4 {
		t.Fatal("down overflow")
	}
	state = Transition(state, Select)
	if state.Screen != Help {
		t.Fatal("help selection")
	}
	state = Transition(state, Back)
	if state.Screen != Home {
		t.Fatal("back")
	}
	state = Transition(state, Quit)
	if !state.Quit {
		t.Fatal("quit")
	}
}

type fakeProvider struct {
	calls int
	value presentationmodel.Overview
	err   error
}

func (f *fakeProvider) Refresh(context.Context) (presentationmodel.Overview, error) {
	f.calls++
	return f.value, f.err
}

func TestRefreshIsExplicitOnceAndRetainsLastGood(t *testing.T) {
	initial := unavailable(t)
	state, _ := NewState(initial, "en", Capabilities{Width: 80, Height: 30})
	provider := &fakeProvider{value: initial}
	var out bytes.Buffer
	if err := Run(context.Background(), strings.NewReader("r\nq\n"), &out, provider, state); err != nil {
		t.Fatal(err)
	}
	if provider.calls != 1 {
		t.Fatalf("calls=%d", provider.calls)
	}
	failed := ApplyRefresh(state, presentationmodel.Overview{}, errors.New("no"))
	if failed.Overview.ID != initial.ID || failed.Diagnostic != "refresh_failed" {
		t.Fatal("failed refresh replaced state")
	}
}

func TestInteractiveSessionRendersInitialOverviewOnceAndClearsOnlyOnRedraw(t *testing.T) {
	initial := unavailable(t)
	state, _ := NewState(initial, "en", Capabilities{Interactive: true, Width: 80, Height: 30})
	var out bytes.Buffer
	if err := Run(context.Background(), strings.NewReader("r\nq\n"), &out, &fakeProvider{value: initial}, state); err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(out.String(), "\x1b[H\x1b[2J")
	if len(parts) != 3 {
		t.Fatalf("redraw boundaries=%d", len(parts)-1)
	}
	if strings.Count(parts[0], text("en", "title")) != 1 {
		t.Fatalf("initial overview count=%d", strings.Count(parts[0], text("en", "title")))
	}
}

func TestTerminalSafetyBoundsAndFallback(t *testing.T) {
	state, _ := NewState(unavailable(t), "en", Capabilities{Width: 40, Height: 8})
	output := Render(state)
	if strings.Contains(output, "\x1b") || len(strings.Split(strings.TrimSuffix(output, "\n"), "\n")) > 8 {
		t.Fatalf("unsafe bounded output: %q", output)
	}
	if ParseAction(strings.Repeat("x", MaxInput+1)) != Unsupported {
		t.Fatal("excess input accepted")
	}
}

func TestCatalogsContainRequiredConsoleTokens(t *testing.T) {
	required := []string{"title", "condition", "attention", "attention_summary", "changes", "alerts", "guardian", "evidence", "recommendation", "navigation", "details", "help", "keys", "refresh_failed", "refresh_invalid", "empty"}
	for locale, catalog := range catalogs {
		for _, token := range required {
			if catalog[token] == "" {
				t.Fatalf("%s missing %s", locale, token)
			}
		}
	}
}

func TestRuntimeFailureDiagnosticsAreLocalized(t *testing.T) {
	for locale, expected := range map[string]map[string]string{
		"en": {"alert_evaluation_failed": "Alert evaluation failed", "notification_planning_failed": "Notification planning failed", "notification_cycle_failed": "Notification delivery failed", "runtime_timeout": "Runtime timed out", "runtime_cancelled": "Runtime was cancelled"},
		"hu": {"alert_evaluation_failed": "A riasztáskiértékelés sikertelen", "notification_planning_failed": "Az értesítéstervezés sikertelen", "notification_cycle_failed": "Az értesítéskézbesítés sikertelen", "runtime_timeout": "A Runtime időtúllépés miatt leállt", "runtime_cancelled": "A Runtime futása megszakadt"},
	} {
		for token, want := range expected {
			if got := text(locale, token); got != want || strings.Contains(got, "[") {
				t.Fatalf("%s %s=%q", locale, token, got)
			}
		}
	}
}

func TestAttentionReductionDisclosureIsLocalizedAndExplicit(t *testing.T) {
	for _, locale := range []string{"en", "hu"} {
		overview := unavailable(t)
		overview.AttentionSummary = &presentationmodel.AttentionSummary{TotalCandidates: 732, Represented: 256, CorrelatedDuplicates: 366, Omitted: 110}
		state := State{Locale: locale, Screen: Home, Capabilities: Capabilities{Width: MaxWidth, Height: MaxHeight}, Overview: overview}
		output := Render(state)
		expected := fmt.Sprintf(text(locale, "attention_summary"), 366, 110)
		if !strings.Contains(output, expected) {
			t.Fatalf("%s did not disclose non-exhaustive attention: %q", locale, output)
		}
	}
}
