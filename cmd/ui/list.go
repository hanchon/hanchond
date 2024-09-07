package ui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/hanchon/hanchond/playground/explorer"
	"github.com/hanchon/hanchond/playground/explorer/database"
)

type block struct {
	text   string
	desc   string
	hash   string
	height int64
}

func (i block) Title() string       { return i.text }
func (i block) Description() string { return i.desc }
func (i block) FilterValue() string { return i.text }

func BDBlockToItem(blocks []database.Block) []list.Item {
	res := make([]list.Item, len(blocks))
	for k := range res {
		res[k] = block{
			text:   fmt.Sprintf("%d", blocks[k].Height),
			desc:   fmt.Sprintf("%s...%s", blocks[k].Hash[0:4], blocks[k].Hash[len(blocks[k].Hash)-5:]),
			height: blocks[k].Height,
			hash:   blocks[k].Hash,
		}
	}
	return res
}

func RenderBlock(b block, client *explorer.Client) string {
	blockData, err := client.Client.GetBlockCosmos(fmt.Sprintf("%d", b.height))
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	data, err := json.MarshalIndent(blockData, "", "  ")
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	cosmosBlock := fmt.Sprintf("# Block %d\n\n## Cosmos Block\n\n```json\n%s\n```", b.height, string(data))

	ethBlock, err := client.Client.GetBlockByNumber(fmt.Sprintf("%d", b.height), true)
	if err != nil {
		return "# Error getting eth block info\n\n" + err.Error()
	}

	data, err = json.MarshalIndent(ethBlock.Result, "", "  ")
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	return cosmosBlock + fmt.Sprintf("\n\n## Ethereum Block\n\n```json\n%s\n```", string(data))
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
		return i.ethHash
	}
	return i.cosmosHash
}

func (i txn) Description() string {
	return i.typeURL
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

func BDTxToItem(txns []database.Transaction) []list.Item {
	res := make([]list.Item, len(txns))
	for k := range res {
		res[k] = txn{
			cosmosHash:  txns[k].Cosmoshash,
			ethHash:     txns[k].Ethhash,
			typeURL:     txns[k].Typeurl,
			sender:      txns[k].Sender,
			blockHeight: int(txns[k].Blockheight),
		}
	}
	return res
}

func RenderTx(b txn, client *explorer.Client) string {
	cosmosTX, err := client.Client.GetCosmosTx(b.cosmosHash)
	if err != nil {
		return "# Error getting cosmos tx\n\n" + err.Error()
	}

	data, err := json.MarshalIndent(cosmosTX, "", "  ")
	if err != nil {
		return "# Error getting cosmos tx\n\n" + err.Error()
	}

	if !strings.Contains(b.typeURL, "ethermint.evm.v1.MsgEthereumTx") {
		return fmt.Sprintf("# Transaction Details\n\n## Cosmos TX:\n- TxHash: %s\n```json\n%s\n```", b.cosmosHash, string(data))
	}

	ethReceipt, err := client.Client.GetTransactionReceipt(b.ethHash)
	if err != nil {
		return "# Error getting eth receipt\n\n" + err.Error()
	}

	ethReceiptString, err := json.MarshalIndent(ethReceipt.Result, "", "  ")
	if err != nil {
		return "# Error getting eth receipt\n\n" + err.Error()
	}

	ethTrace, err := client.Client.GetTransactionTrace(b.ethHash)
	if err != nil {
		return "# Error getting eth trace\n\n" + err.Error()
	}

	ethTraceString, err := json.MarshalIndent(ethTrace.Result, "", "  ")
	if err != nil {
		return "# Error getting eth trace\n\n" + err.Error()
	}

	return fmt.Sprintf("# Transaction Details\n\n## Ethereum Transaction:\n- TxHash: %s\n ### Receipt:\n```json\n%s\n```\n### Trace:\n```json\n%s\n```\n## Cosmos Transaction:\n- TxHash: %s\n```json\n%s\n```",
		b.ethHash,
		string(ethReceiptString),
		string(ethTraceString),
		b.cosmosHash,
		string(data),
	)
}
