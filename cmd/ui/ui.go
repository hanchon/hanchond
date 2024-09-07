package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/hanchon/hanchond/cmd/ui/explorer"
	explorerClient "github.com/hanchon/hanchond/playground/explorer"
)

var mdRendered *glamour.TermRenderer

var basicStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center).
	AlignHorizontal(lipgloss.Center)

type model struct {
	width  int
	height int

	activeList int
	lists      []list.Model
	viewport   viewport.Model

	mdValues string

	client         *explorerClient.Client
	startingHeight int64

	resolutionError bool
}

func (m model) Init() tea.Cmd {
	return nil
}

type tickMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		// If it is already running it will return a no-op
		go m.client.ProcessMissingBlocks(m.startingHeight)
		b, t, err := m.client.DB.GetDisplayInfo(50)
		if err == nil {
			m.lists[0].SetItems(BDBlockToItem(b))
			m.lists[1].SetItems(BDTxToItem(t))
		}
		return m, tickCmd()
	case tea.WindowSizeMsg:
		m.height = msg.(tea.WindowSizeMsg).Height
		m.width = msg.(tea.WindowSizeMsg).Width - 2
		m.mdValues = fmt.Sprintf("%d %d", msg.(tea.WindowSizeMsg).Height, msg.(tea.WindowSizeMsg).Width)
		basicStyle = basicStyle.
			Width(m.width - 2).
			Height(m.height)

		if m.height < 48 || m.width+2 < 190 {
			m.resolutionError = true
		} else {
			m.resolutionError = false
		}
		return m, tickCmd()
	case tea.KeyMsg:
		key := msg.(tea.KeyMsg).String()
		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}
		if key == "tab" {
			m.activeList = (m.activeList + 1) % 3
			return m, nil
		}

		if key == "shift+tab" {
			m.activeList = (m.activeList - 1) % 3
			return m, nil
		}
		if key == "enter" {
			switch m.activeList {
			case 0:
				selectedItem := m.lists[0].SelectedItem().(block)
				info, _ := mdRendered.Render(RenderBlock(selectedItem, m.client))
				m.viewport.SetContent(info)
				m.viewport.Height = 23
				m.viewport.Width = 78
				return m, nil
			case 1:
				selectedItem := m.lists[1].SelectedItem().(txn)
				info, _ := mdRendered.Render(RenderTx(selectedItem, m.client))
				m.viewport.SetContent(info)
				m.viewport.Height = 23
				m.viewport.Width = 78
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
	if m.resolutionError {
		return lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Foreground(explorer.ColorHighPink).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render("Your resolution is too low, please reduce the zoom to display the dashboard")
	}

	value := lipgloss.JoinVertical(
		lipgloss.Top,
		explorer.Header(m.width-4),
		explorer.ChainHeightFrame(m.width-4, m.client.NetworkHeight, m.client.DBHeight),
		explorer.BotContainer(m.width-4, m.lists[0].View(), m.lists[1].View(), m.viewport.View(), m.activeList),
	)

	return basicStyle.Render(value)
}

func CreateExplorerTUI(startHeight int, client *explorerClient.Client) *tea.Program {
	mdRendered, _ = glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(78),
	)

	m := model{
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
