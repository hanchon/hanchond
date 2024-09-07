package explorerui

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/hanchon/hanchond/playground/explorer"
	"github.com/hanchon/hanchond/playground/explorer/database"
)

type Block struct {
	text   string
	desc   string
	hash   string
	height int64
}

func (i Block) Title() string       { return i.text }
func (i Block) Description() string { return i.desc }
func (i Block) FilterValue() string { return i.text }

func BDBlockToItem(blocks []database.Block) []list.Item {
	res := make([]list.Item, len(blocks))
	for k := range res {
		res[k] = Block{
			text:   fmt.Sprintf("%d", blocks[k].Height),
			desc:   fmt.Sprintf("%s...%s", blocks[k].Hash[0:4], blocks[k].Hash[len(blocks[k].Hash)-5:]),
			height: blocks[k].Height,
			hash:   blocks[k].Hash,
		}
	}
	return res
}

func RenderBlock(b Block, client *explorer.Client) string {
	blockData, err := client.Client.GetBlockCosmos(fmt.Sprintf("%d", b.height))
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	data, err := json.MarshalIndent(blockData, "", "  ")
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	cosmosBlock := fmt.Sprintf("# Block %d\n\n## Cosmos Block\n\n```json\n%s\n```", b.height, processJSON(string(data)))

	ethBlock, err := client.Client.GetBlockByNumber(fmt.Sprintf("%d", b.height), true)
	if err != nil {
		return "# Error getting eth block info\n\n" + err.Error()
	}

	data, err = json.MarshalIndent(ethBlock.Result, "", "  ")
	if err != nil {
		return "# Error getting block info\n\n" + err.Error()
	}

	return cosmosBlock + fmt.Sprintf("\n\n## Ethereum Block\n\n```json\n%s\n```", processJSON(string(data)))
}
