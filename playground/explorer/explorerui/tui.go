package explorerui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"

	explorerClient "github.com/hanchon/hanchond/playground/explorer"
)

func CreateExplorerTUI(startHeight int, client *explorerClient.Client) *tea.Program {
	mdRendered, _ = glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(78),
	)

	m := explorerModel{
		client:   client,
		mdValues: "",
	}

	list1 := list.New([]list.Item{}, list.NewDefaultDelegate(), 20, 14)
	list1.Title = "Latest Blocks"
	list1.SetWidth(20)
	list1.SetHeight(23)
	list1.Styles.TitleBar.Align(lipgloss.Center)

	list2 := list.New([]list.Item{}, list.NewDefaultDelegate(), 20, 14)
	list2.Title = "Latest Transactions"
	list2.SetWidth(80)
	list2.SetHeight(23)

	m.lists = append(m.lists, list1)
	m.lists = append(m.lists, list2)

	m.viewport = viewport.New(78, 23)
	m.startingHeight = int64(startHeight)

	go client.ProcessMissingBlocks(int64(startHeight))

	return tea.NewProgram(m)
}
