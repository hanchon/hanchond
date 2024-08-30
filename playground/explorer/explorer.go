package explorer

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/protoencoder/codec"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/explorer/database"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Client struct {
	web3Endpoint   string
	cosmosEndpoint string

	client *requester.Client

	ctx context.Context

	db *Database
}

func NewLocalExplorerClient(web3Port, cosmosPort int, homeFolder string) *Client {
	queries, err := database.InitExplorerDatabase(context.Background(), homeFolder+"/explorer.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	e := &Client{
		web3Endpoint:   fmt.Sprintf("http://localhost:%d", web3Port),
		cosmosEndpoint: fmt.Sprintf("http://localhost:%d", cosmosPort),
		ctx:            context.Background(),
	}
	e.db = NewDatabase(e.ctx, queries)
	e.client = requester.NewClient().WithUnsecureWeb3Endpoint(e.web3Endpoint).WithUnsecureRestEndpoint(e.cosmosEndpoint)

	return e
}

func (e *Client) ProcessBlocks() {
	// TODO: Delete the last bock in case the last execution was not completed
	block, err := e.db.GetLatestBlock()

	initBlockHeight := int64(0)
	if err == nil {
		initBlockHeight = block.Height + 1
	}

	networkHeight, err := e.client.GetBlockNumber()
	if err != nil {
		fmt.Println("error getting latest block")
		os.Exit(1)
	}

	// TODO: if network height < current sleep and retry
	if networkHeight > initBlockHeight {
		height := 2332
		blockData, err := e.client.GetBlockCosmos(fmt.Sprintf("0x%x", height))
		if err != nil {
			panic(err)
		}
		blockHash, err := converter.Base64ToHexString(blockData.BlockID.Hash)
		if err != nil {
			panic(err)
		}

		data := NewBlock(int64(height), int64(len(blockData.Block.Data.Txs)), blockHash)

		for i, txBase64 := range blockData.Block.Data.Txs {
			tx, err := codec.Base64ToTx(txBase64)
			if err != nil {
				panic(err)
			}

			sender := sdk.AccAddress(tx.AuthInfo.GetSignerInfos()[0].PublicKey.Value).String()
			if len(tx.Body.Messages) == 0 {
				panic(err)
			}

			typeURL := tx.Body.Messages[0].TypeUrl
			cosmosTxHash, err := converter.GenerateCosmosTxHashWithBase64(txBase64)
			if err != nil {
				panic(err)
			}

			ethTxHash := ""
			ethTx, from, err := codec.ConvertEvmosTxToEthTx(txBase64)
			if err == nil {
				ethTxHash = ethTx.Hash().Hex()
				sender = from.String()
			}

			sender, _ = converter.Bech32ToHex(sender)
			data.AddTransaction(i, cosmosTxHash, ethTxHash, typeURL, sender)
		}
		fmt.Println(data)
		os.Exit(0)
	}
}
