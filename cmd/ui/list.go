package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type block struct {
	text string
	desc string
}

func (i block) Title() string       { return i.text }
func (i block) Description() string { return i.desc }
func (i block) FilterValue() string { return i.text }

var items1 = []list.Item{
	block{text: "Block: 10000000", desc: "1A25...9BCB"},
	block{text: "Block: 10000001", desc: "2A25...9BCB"},
	block{text: "Block: 10000002", desc: "3A25...9BCB"},
}

type txn struct {
	cosmosHash  string
	ethHash     string
	typeURL     string
	sender      string
	blockHeight int
}

func (i txn) Title() string {
	if i.ethHash != "" {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
		return style.Render(i.ethHash)
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("35")).Render(i.cosmosHash)
}

func (i txn) Description() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("75"))
	return style.Render(i.typeURL)
}

// TODO: this should filter by everything
func (i txn) FilterValue() string { return i.ethHash }

var style = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))

var items2 = []list.Item{
	txn{
		ethHash:    "0x61b7f582cfe2ee3b9d31dcbf99e5036b1c68713ede8ce7ed13930f2e02470588",
		cosmosHash: "0x61b7...0588",
		typeURL:    "MsgEthereum",
	},
	txn{
		ethHash:    "",
		cosmosHash: "0x71b7...0588",
		typeURL:    "MsgVote",
	},
}
