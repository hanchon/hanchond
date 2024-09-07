package explorerui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/hanchon/hanchond/playground/explorer"
	"github.com/hanchon/hanchond/playground/explorer/database"
)

type Txn struct {
	cosmosHash  string
	ethHash     string
	typeURL     string
	sender      string
	blockHeight int
}

func (i Txn) Title() string {
	if i.ethHash != "" {
		return i.ethHash
	}
	return i.cosmosHash
}

func (i Txn) Description() string {
	return i.typeURL
}

// TODO: this should filter by everything
func (i Txn) FilterValue() string { return strings.ToLower(i.typeURL) }

func BDTxToItem(txns []database.Transaction) []list.Item {
	res := make([]list.Item, len(txns))
	for k := range res {
		res[k] = Txn{
			cosmosHash:  txns[k].Cosmoshash,
			ethHash:     txns[k].Ethhash,
			typeURL:     txns[k].Typeurl,
			sender:      txns[k].Sender,
			blockHeight: int(txns[k].Blockheight),
		}
	}
	return res
}

func RenderTx(b Txn, client *explorer.Client) string {
	cosmosTX, err := client.Client.GetCosmosTx(b.cosmosHash)
	if err != nil {
		return "# Error getting cosmos tx\n\n" + err.Error()
	}

	data, err := json.MarshalIndent(cosmosTX, "", "  ")
	if err != nil {
		return "# Error getting cosmos tx\n\n" + err.Error()
	}

	if !strings.Contains(b.typeURL, "ethermint.evm.v1.MsgEthereumTx") {
		return fmt.Sprintf("# Transaction Details\n\n## Cosmos TX:\n- Status: %v\n- TxHash: %s\n```json\n%s\n```", cosmosTX.TxResponse.Code == 0, b.cosmosHash, string(data))
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

	return fmt.Sprintf("# Transaction Details\n\n## Ethereum Transaction:\n- Status: %v\n- TxHash: %s\n ### Receipt:\n```json\n%s\n```\n### Trace:\n```json\n%s\n```\n## Cosmos Transaction:\n- Status: %v\n- TxHash: %s\n```json\n%s\n```",
		ethReceipt.Result.Status == "0x1",
		b.ethHash,
		processJSON(string(ethReceiptString)),
		processJSON(string(ethTraceString)),
		cosmosTX.TxResponse.Code == 0,
		b.cosmosHash,
		processJSON(string(data)),
	)
}
