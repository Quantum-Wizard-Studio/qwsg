package operatorconsole

import (
	"fmt"
	"strings"
	"unicode"

	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

func Render(state State) string {
	lines := []string{text(state.Locale, "title"), strings.Repeat("=", 22)}
	view := state.Overview
	switch state.Screen {
	case Home:
		lines = append(lines,
			pair(state, "condition", string(view.Condition)), pair(state, "attention", string(view.Attention)),
		)
		lines = appendAttentionSummary(lines, state)
		lines = append(lines, fmt.Sprintf("%s: %d", text(state.Locale, "changes"), view.Summary.Changes),
			fmt.Sprintf("%s: %d", text(state.Locale, "alerts"), view.Summary.ActiveAlerts),
			pair(state, "guardian", string(view.Guardian)),
			fmt.Sprintf("%s: %s / %s", text(state.Locale, "evidence"), text(state.Locale, string(view.Freshness)), text(state.Locale, string(view.Completeness))),
			fmt.Sprintf("%s: %s", text(state.Locale, "recommendation"), recommendation(state)), "", text(state.Locale, "navigation"),
		)
		for i, token := range []string{"attention", "changes", "guardian", "details", "help"} {
			lines = append(lines, marker(i == state.Selection)+text(state.Locale, token))
		}
	case Attention:
		lines = append(lines, text(state.Locale, "attention"))
		if len(view.AttentionItems) == 0 {
			lines = append(lines, text(state.Locale, "empty"))
		}
		for i, item := range view.AttentionItems {
			lines = append(lines, marker(i == state.Selection)+text(state.Locale, string(item.Severity))+": "+text(state.Locale, item.TitleToken))
		}
		lines = appendAttentionSummary(lines, state)
	case Changes:
		lines = append(lines, text(state.Locale, "changes"))
		if len(view.Changes) == 0 {
			lines = append(lines, text(state.Locale, "empty"))
		}
		for i, item := range view.Changes {
			lines = append(lines, fmt.Sprintf("%s%s: +%d -%d ~%d", marker(i == state.Selection), safe(item.Category), item.Added, item.Removed, item.Modified))
		}
	case Guardian:
		lines = append(lines, pair(state, "guardian", string(view.Guardian)), fmt.Sprintf("%s: %s / %s", text(state.Locale, "evidence"), text(state.Locale, string(view.Freshness)), text(state.Locale, string(view.Completeness))))
	case Details:
		lines = append(lines, text(state.Locale, "details"))
		if len(view.Sources) == 0 {
			lines = append(lines, text(state.Locale, "empty"))
		}
		for i, item := range view.Sources {
			lines = append(lines, sourceLine(state, i, item))
		}
	case Help:
		lines = append(lines, text(state.Locale, "keys"))
	}
	if state.Diagnostic != "" {
		lines = append(lines, "", text(state.Locale, state.Diagnostic))
	}
	lines = append(lines, "", text(state.Locale, "keys"))
	for i := range lines {
		lines[i] = truncate(safe(lines[i]), state.Capabilities.Width)
	}
	if len(lines) > state.Capabilities.Height {
		lines = lines[:state.Capabilities.Height]
	}
	return strings.Join(lines, "\n") + "\n"
}

func appendAttentionSummary(lines []string, state State) []string {
	if summary := state.Overview.AttentionSummary; summary != nil {
		lines = append(lines, fmt.Sprintf(text(state.Locale, "attention_summary"), summary.CorrelatedDuplicates, summary.Omitted))
	}
	return lines
}

func RenderPlain(overview presentationmodel.Overview, locale string) (string, error) {
	state, err := NewState(overview, locale, Capabilities{Width: MaxWidth, Height: 20})
	if err != nil {
		return "", err
	}
	return Render(state), nil
}

func pair(state State, label, value string) string {
	return text(state.Locale, label) + ": " + text(state.Locale, value)
}
func marker(selected bool) string {
	if selected {
		return "> "
	}
	return "  "
}
func recommendation(state State) string {
	if len(state.Overview.Recommendations) == 0 {
		return text(state.Locale, "none")
	}
	return text(state.Locale, string(state.Overview.Recommendations[0].Token))
}
func sourceLine(state State, index int, source presentationmodel.SourceReference) string {
	return fmt.Sprintf("%s%s: %s %s %s", marker(index == state.Selection), text(state.Locale, "source"), safe(source.Kind), safe(source.RecordID), source.ObservedAt.UTC().Format("2006-01-02T15:04:05Z"))
}
func safe(value string) string {
	var b strings.Builder
	for _, r := range value {
		if unicode.IsControl(r) || r == 0x7f {
			b.WriteRune('�')
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
func truncate(value string, width int) string {
	runes := []rune(value)
	if len(runes) <= width {
		return value
	}
	if width < 2 {
		return ""
	}
	return string(runes[:width-1]) + "…"
}
