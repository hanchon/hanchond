package explorerui

import "github.com/charmbracelet/lipgloss"

var headerStyle = lipgloss.NewStyle().
	Foreground(ColorHighPink).
	Height(1).
	Padding(1, 1).
	Align(lipgloss.Center).
	Border(noClosingBot, true).
	Bold(true).
	BorderForeground(ColorHighlight)

func Header(width int) string {
	return headerStyle.Width(width).Render("Block Explorer")
}
