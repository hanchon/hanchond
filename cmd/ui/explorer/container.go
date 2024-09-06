package explorer

import (
	"github.com/charmbracelet/lipgloss"
)

var botContainerStyle = lipgloss.NewStyle().
	Height(25).
	AlignVertical(lipgloss.Center).
	Align(lipgloss.Center).
	Border(lipgloss.NormalBorder(), true).
	BorderTop(false).
	BorderForeground(ColorHighlight)

var blocksFrame = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Height(24)
var txFrame = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Height(24)
var infoFrame = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Height(24)

func BotContainer(width int, list1, list2 string, infoText string) string {
	blocks := blocksFrame.Width(20)
	transactions := txFrame.Width((width-20)/2 - 2)
	info := infoFrame.Width((width-20)/2 - 4)

	temp := lipgloss.JoinHorizontal(
		lipgloss.Center,
		blocks.Render(list1),
		transactions.Render(list2),
		info.Render(infoText),
	)

	return botContainerStyle.
		Width(width).
		Render(temp)

}
