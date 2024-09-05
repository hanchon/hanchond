package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hanchon/hanchond/cmd/ui/explorer"
	// "github.com/charmbracelet/log"
)

type model struct {
	width  int
	height int

	activeList int // Index of the currently active list
	lists      []list.Model
}

func (m model) Init() tea.Cmd {
	// No initial command
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.(tea.WindowSizeMsg).Height
		m.width = msg.(tea.WindowSizeMsg).Width - 2
	case tea.KeyMsg:
		// Exit on any key press
		return m, tea.Quit
	}
	return m, nil
}
func (m model) View() string {
	// log.Info(m.height)
	// log.Info(m.width)

	value := lipgloss.JoinVertical(
		lipgloss.Top,
		explorer.Header(m.width-4),
		explorer.ChainHeightFrame(m.width-4, 10000, 5000),
		explorer.BotContainer(m.width-4, m.lists[0].View(), m.lists[1].View()),
	)

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Width(m.width - 2).
		Height(m.height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	return style.Render(value)
}

func CreateExplorerTUI() *tea.Program {
	m := model{}

	list1 := list.New(items1, list.NewDefaultDelegate(), 20, 14)
	list1.Title = "Latest Blocks"
	list1.SetWidth(20)
	list1.SetHeight(23)
	list1.Styles.TitleBar.Align(lipgloss.Center)

	list2 := list.New(items2, list.NewDefaultDelegate(), 20, 14)
	list2.Title = "Latest Transactions"
	list2.SetWidth(80)
	list2.SetHeight(23)
	m.lists = append(m.lists, list1)
	m.lists = append(m.lists, list2)
	return tea.NewProgram(m)
}
