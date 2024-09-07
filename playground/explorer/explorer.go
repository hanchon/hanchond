package explorer

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/protoencoder/codec"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/explorer/database"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Client struct {
	mutex          sync.Mutex
	web3Endpoint   string
	cosmosEndpoint string

	Client *requester.Client

	ctx context.Context

	DB *Database

	NetworkHeight int
	DBHeight      int
}

func NewLocalExplorerClient(web3Port, cosmosPort int, homeFolder string) *Client {
	db, queries, err := database.InitExplorerDatabase(context.Background(), homeFolder+"/explorer.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	c := &Client{
		mutex:          sync.Mutex{},
		web3Endpoint:   fmt.Sprintf("http://localhost:%d", web3Port),
		cosmosEndpoint: fmt.Sprintf("http://localhost:%d", cosmosPort),
		ctx:            context.Background(),
		NetworkHeight:  0,
		DBHeight:       0,
	}
	c.DB = NewDatabase(c.ctx, db, queries)
	c.Client = requester.NewClient().WithUnsecureWeb3Endpoint(c.web3Endpoint).WithUnsecureRestEndpoint(c.cosmosEndpoint)

	return c
}

// ProcessMissingBlocks process up to 500 blocks at the time
func (c *Client) ProcessMissingBlocks(startBlock int64) error {
	if !c.mutex.TryLock() {
		return nil
	}
	defer c.mutex.Unlock()
	blocksData := []Block{}
	nextBlockToIndex := startBlock

	block, err := c.DB.GetLatestBlock()
	if err == nil {
		if block.Height > startBlock {
			nextBlockToIndex = block.Height + 1
		}
	}

	networkHeight, err := c.Client.GetBlockNumber()
	if err != nil {
		return fmt.Errorf("error getting latest block: %s", err.Error())
	}

	c.DBHeight = int(block.Height)
	c.NetworkHeight = int(networkHeight)

	if networkHeight < nextBlockToIndex {
		// We are up to date
		return nil
	}

	if networkHeight-500 > nextBlockToIndex {
		// Batch no more than 500 blocks
		networkHeight = nextBlockToIndex + 500
	}

	for networkHeight >= nextBlockToIndex {
		blockData, err := c.Client.GetBlockCosmos(fmt.Sprintf("0x%x", nextBlockToIndex))
		if err != nil {
			return fmt.Errorf("error getting cosmos block: %s", err.Error())
		}
		blockHash, err := converter.Base64ToHexString(blockData.BlockID.Hash)
		if err != nil {
			return fmt.Errorf("error getting cosmos block hash: %s", err.Error())
		}

		data := NewBlock(nextBlockToIndex, int64(len(blockData.Block.Data.Txs)), blockHash)

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
		nextBlockToIndex++
		blocksData = append(blocksData, *data)
	}

	// Save it to the database
	err = c.DB.AddBlocks(blocksData)
	if err != nil {
		return err
	}

	c.DBHeight = int(nextBlockToIndex)
	return nil
}
