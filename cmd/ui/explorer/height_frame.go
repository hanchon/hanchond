package explorer

import (
	"strconv"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

var chainFrameStyle = lipgloss.NewStyle().
	Foreground(ColorPurple).
	Height(1).
	Align(lipgloss.Center).
	BorderForeground(ColorHighlight)

var chainFramStyleText = lipgloss.NewStyle().
	Foreground(ColorPurple).
	Height(1).
	Align(lipgloss.Left)

var graphModel = progress.New(progress.WithDefaultGradient())
var graphModelWidth = graphModel.Width

var heightFrameStyle = lipgloss.NewStyle().
	Padding(0, 3).
	Border(lipgloss.NormalBorder(), false, true, false, false).
	BorderForeground(ColorHighlight).
	Width(graphModelWidth).
	AlignHorizontal(lipgloss.Right)

var graphFrameStyle = lipgloss.NewStyle().
	PaddingTop(1).
	PaddingLeft(1).
	Height(4)

func ChainHeightFrame(width, currentHeightInt, indexerHeightInt int) string {
	currentHeight := strconv.Itoa(currentHeightInt)
	bHeight := lipgloss.JoinHorizontal(
		lipgloss.Left,
		chainFramStyleText.Render("۩ Current Height: "),
		chainFramStyleText.Render(currentHeight),
	)

	indexerHeight := strconv.Itoa(indexerHeightInt)
	for range len(currentHeight) - len(indexerHeight) {
		indexerHeight = " " + indexerHeight
	}

	iHeight := lipgloss.JoinHorizontal(
		lipgloss.Left,
		chainFrameStyle.Render("֍ Indexer Height: "),
		chainFrameStyle.Render(indexerHeight),
	)

	heightFrame := heightFrameStyle.
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				chainFrameStyle.Bold(true).Foreground(ColorLowPink).Render("Blocks Info"),
				bHeight,
				iHeight,
			),
		)

	graphFrame := graphFrameStyle.Render(
		graphModel.ViewAs(float64(indexerHeightInt) / float64(currentHeightInt)),
	)

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		heightFrame,
		graphFrame,
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		chainFrameStyle.
			Border(noClosingTopBot, false, true, false).
			Foreground(lipgloss.Color("32")).
			PaddingBottom(1).
			Bold(true).
			Width(width).
			Render("Chain Height"),
		chainFrameStyle.
			Border(noClosingTopBot, false, true, true).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Width(width).
			Render(content),
	)
}
