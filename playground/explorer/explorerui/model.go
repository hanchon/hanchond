package explorerui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	explorerClient "github.com/hanchon/hanchond/playground/explorer"
)

var mdRendered *glamour.TermRenderer

var basicStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center).
	AlignHorizontal(lipgloss.Center)

type explorerModel struct {
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

func (m explorerModel) Init() tea.Cmd {
	return nil
}

type tickMsg struct{}

func indexerTickerCmd() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(_ time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m explorerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tickMsg:
		// If it is already running it will return a no-op
		go m.client.ProcessMissingBlocks(m.startingHeight) //nolint: errcheck
		b, t, err := m.client.DB.GetDisplayInfo(50)
		if err == nil {
			m.lists[0].SetItems(BDBlockToItem(b))
			m.lists[1].SetItems(BDTxToItem(t))
		}
		return m, indexerTickerCmd()
	case tea.WindowSizeMsg:
		m.height = v.Height
		m.width = v.Width - 2
		m.mdValues = fmt.Sprintf("%d %d", msg.(tea.WindowSizeMsg).Height, msg.(tea.WindowSizeMsg).Width)
		basicStyle = basicStyle.
			Width(m.width - 2).
			Height(m.height)

		if m.height < 48 || m.width+2 < 190 {
			m.resolutionError = true
		} else {
			m.resolutionError = false
		}
		return m, indexerTickerCmd()
	case tea.KeyMsg:
		key := v.String()
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
				selectedItem := m.lists[0].SelectedItem().(Block)
				info, _ := mdRendered.Render(RenderBlock(selectedItem, m.client))
				m.viewport.SetContent(info)
				_ = m.viewport.GotoTop()
				return m, nil
			case 1:
				selectedItem := m.lists[1].SelectedItem().(Txn)
				info, _ := mdRendered.Render(RenderTx(selectedItem, m.client))
				m.viewport.SetContent(info)
				_ = m.viewport.GotoTop()
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

func (m explorerModel) View() string {
	if m.resolutionError {
		return lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Foreground(ColorHighPink).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render("Your resolution is too low, please reduce the zoom to display the dashboard")
	}

	value := lipgloss.JoinVertical(
		lipgloss.Top,
		Header(m.width-4),
		ChainHeightFrame(m.width-4, m.client.NetworkHeight, m.client.DBHeight),
		BotContainer(m.width-4, m.lists[0].View(), m.lists[1].View(), m.viewport.View(), m.activeList),
	)

	return basicStyle.Render(value)
}
