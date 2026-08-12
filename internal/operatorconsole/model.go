// Package operatorconsole renders and navigates validated operator overviews.
// It is presentation-only and owns no engineering or operational decisions.
package operatorconsole

import (
	"context"
	"fmt"
	"strings"

	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

const (
	ModelVersion = "1.0"
	MaxInput     = 64
	MaxWidth     = 240
	MaxHeight    = 100
)

type Screen string

const (
	Home      Screen = "home"
	Attention Screen = "attention"
	Changes   Screen = "changes"
	Guardian  Screen = "guardian"
	Details   Screen = "details"
	Help      Screen = "help"
)

type Action string

const (
	Up          Action = "up"
	Down        Action = "down"
	Select      Action = "select"
	Back        Action = "back"
	Refresh     Action = "refresh"
	ShowHelp    Action = "help"
	Quit        Action = "quit"
	Unsupported Action = "unsupported"
)

type Capabilities struct {
	Interactive bool
	ANSI        bool
	Color       bool
	Width       int
	Height      int
}

type OverviewProvider interface {
	Refresh(context.Context) (presentationmodel.Overview, error)
}

type State struct {
	Version      string
	Screen       Screen
	Selection    int
	Offset       int
	Locale       string
	Capabilities Capabilities
	Overview     presentationmodel.Overview
	Diagnostic   string
	Refreshes    uint64
	Quit         bool
}

func NewState(overview presentationmodel.Overview, locale string, capabilities Capabilities) (State, error) {
	if err := presentationmodel.Validate(overview); err != nil {
		return State{}, fmt.Errorf("invalid operator overview: %w", err)
	}
	if _, ok := catalogs[locale]; !ok {
		locale = "en"
	}
	capabilities.Width = bounded(capabilities.Width, 40, MaxWidth)
	capabilities.Height = bounded(capabilities.Height, 8, MaxHeight)
	return State{Version: ModelVersion, Screen: Home, Locale: locale, Capabilities: capabilities, Overview: overview}, nil
}

func ParseAction(value string) Action {
	if len(value) > MaxInput {
		return Unsupported
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "k", "up":
		return Up
	case "j", "down":
		return Down
	case "", "enter", "select":
		return Select
	case "b", "back":
		return Back
	case "r", "refresh":
		return Refresh
	case "h", "?", "help":
		return ShowHelp
	case "q", "quit":
		return Quit
	default:
		return Unsupported
	}
}

func Transition(state State, action Action) State {
	next := state
	switch action {
	case Up:
		if next.Selection > 0 {
			next.Selection--
		}
	case Down:
		if next.Selection+1 < itemCount(next) {
			next.Selection++
		}
	case Select:
		if next.Screen == Home {
			next.Screen = []Screen{Attention, Changes, Guardian, Details, Help}[next.Selection]
			next.Selection = 0
		}
	case Back:
		next.Screen, next.Selection, next.Offset = Home, 0, 0
	case ShowHelp:
		next.Screen, next.Selection = Help, 0
	case Quit:
		next.Quit = true
	}
	return normalize(next)
}

func ApplyRefresh(state State, overview presentationmodel.Overview, err error) State {
	state.Refreshes++
	if err != nil {
		state.Diagnostic = "refresh_failed"
		return state
	}
	if validation := presentationmodel.Validate(overview); validation != nil {
		state.Diagnostic = "refresh_invalid"
		return state
	}
	state.Overview, state.Diagnostic = overview, ""
	return normalize(state)
}

func normalize(state State) State {
	count := itemCount(state)
	if count == 0 {
		state.Selection, state.Offset = 0, 0
	} else if state.Selection >= count {
		state.Selection = count - 1
	}
	if state.Selection < 0 {
		state.Selection = 0
	}
	return state
}

func itemCount(state State) int {
	switch state.Screen {
	case Home:
		return 5
	case Attention:
		return len(state.Overview.AttentionItems)
	case Changes:
		return len(state.Overview.Changes)
	case Details:
		return len(state.Overview.Sources)
	default:
		return 1
	}
}

func bounded(value, minimum, maximum int) int {
	if value < minimum {
		return minimum
	}
	if value > maximum {
		return maximum
	}
	return value
}
