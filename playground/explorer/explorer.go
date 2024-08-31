package explorer

import (
	"context"
	"fmt"
	"log"

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
	db, queries, err := database.InitExplorerDatabase(context.Background(), homeFolder+"/explorer.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	c := &Client{
		web3Endpoint:   fmt.Sprintf("http://localhost:%d", web3Port),
		cosmosEndpoint: fmt.Sprintf("http://localhost:%d", cosmosPort),
		ctx:            context.Background(),
	}
	c.db = NewDatabase(c.ctx, db, queries)
	c.client = requester.NewClient().WithUnsecureWeb3Endpoint(c.web3Endpoint).WithUnsecureRestEndpoint(c.cosmosEndpoint)

	return c
}

func (c *Client) ProcessMissingBlocks(startBlock int64) error {
	blocksData := []Block{}
	currentBlockDB := startBlock

	// TODO: Delete the last bock in case the last execution was not completed
	block, err := c.db.GetLatestBlock()
	if err == nil {
		if block.Height > startBlock {
			currentBlockDB = block.Height + 1
		}
	}

	networkHeight, err := c.client.GetBlockNumber()
	if err != nil {
		return fmt.Errorf("error getting latest block: %s", err.Error())
	}

	// TODO: if network height < current sleep and retry
	for networkHeight > currentBlockDB {
		blockData, err := c.client.GetBlockCosmos(fmt.Sprintf("0x%x", currentBlockDB))
		if err != nil {
			return fmt.Errorf("error getting cosmos block: %s", err.Error())
		}
		blockHash, err := converter.Base64ToHexString(blockData.BlockID.Hash)
		if err != nil {
			return fmt.Errorf("error getting cosmos block hash: %s", err.Error())
		}

		data := NewBlock(currentBlockDB, int64(len(blockData.Block.Data.Txs)), blockHash)

		for i, txBase64 := range blockData.Block.Data.Txs {
			tx, err := codec.Base64ToTx(txBase64)
			if err != nil {
				return fmt.Errorf("error decoding cosmos tx: %s", err.Error())
			}

			if len(tx.Body.Messages) == 0 {
				return fmt.Errorf("error decoding cosmos tx, no messages")
			}

			typeURL := tx.Body.Messages[0].TypeUrl
			cosmosTxHash, err := converter.GenerateCosmosTxHashWithBase64(txBase64)
			if err != nil {
				return fmt.Errorf("error generating cosmos tx hash: %s", err.Error())
			}

			sender := ""
			ethTxHash := ""
			ethTx, from, err := codec.ConvertEvmosTxToEthTx(txBase64)
			if err == nil {
				// Eth Transaction
				ethTxHash = ethTx.Hash().Hex()
				sender = from.String()
			} else if len(tx.AuthInfo.GetSignerInfos()) != 0 {
				// If the transaction was not an Eth Transaction, the sender is in the cosmos signer info
				sender = sdk.AccAddress(tx.AuthInfo.GetSignerInfos()[0].PublicKey.Value).String()
			}

			sender, err = converter.Bech32ToHex(sender)
			if err != nil {
				return fmt.Errorf("error converting sender address: %s", err.Error())
			}
			data.AddTransaction(i, cosmosTxHash, ethTxHash, typeURL, sender)
		}
		currentBlockDB++
		blocksData = append(blocksData, *data)
	}

	// Save it to the database
	return c.db.AddBlocks(blocksData)
}
