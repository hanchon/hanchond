package explorerui

import "github.com/charmbracelet/lipgloss"

var ColorHighlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

var (
	ColorLowPink  = lipgloss.Color("205")
	ColorHighPink = lipgloss.Color("201")
	ColorPurple   = lipgloss.Color("111")
)

var noClosingBot = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "│",
	BottomRight: "│",
}

var noClosingTopBot = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "│",
	TopRight:    "│",
	BottomLeft:  "│",
	BottomRight: "│",
}
