package explorer

import "github.com/charmbracelet/lipgloss"

var ColorHighlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

var ColorLowPink = lipgloss.Color("205")
var ColorHighPink = lipgloss.Color("201")
var ColorPurple = lipgloss.Color("111")

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
