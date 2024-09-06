package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/hanchon/hanchond/cmd/ui/explorer"
	// "github.com/charmbracelet/log"
)

var mdValues = ``

type model struct {
	width  int
	height int

	activeList int
	lists      []list.Model
	viewport   viewport.Model
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
		key := msg.(tea.KeyMsg).String()
		if key == "ctrl+c" {
			return m, tea.Quit
		}
		if key == "tab" {
			m.activeList = (m.activeList + 1) % 3
			return m, nil
		}
		if key == "enter" {
			switch m.activeList {
			case 0:
				selectedItem := m.lists[0].SelectedItem()
				mdValues = selectedItem.(block).text
				return m, nil
			case 1:
				selectedItem := m.lists[1].SelectedItem()
				mdValues = selectedItem.(txn).ethHash
				return m, nil
			}
		}
	}
	var cmd tea.Cmd
	switch m.activeList {
	case 0:
		m.lists[0], cmd = m.lists[0].Update(msg)
		return m, cmd
	case 1:
		m.lists[1], cmd = m.lists[1].Update(msg)
		return m, cmd
	case 2:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}
func (m model) View() string {
	// log.Info(m.height)
	// log.Info(m.width)

	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(78),
	)
	info, _ := r.Render(mdValues)
	m.viewport.SetContent(info)

	value := lipgloss.JoinVertical(
		lipgloss.Top,
		explorer.Header(m.width-4),
		explorer.ChainHeightFrame(m.width-4, 10000, 5000),
		explorer.BotContainer(m.width-4, m.lists[0].View(), m.lists[1].View(), m.viewport.View(), m.activeList),
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

	m.viewport = viewport.New(78, 23)

	return tea.NewProgram(m)
}
