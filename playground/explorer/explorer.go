package explorer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/protoencoder/codec"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/explorer/database"

	evmtypes "github.com/evmos/evmos/v18/x/evm/types"
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
		networkHeight = 474
		a, err := e.client.GetBlockCosmos(fmt.Sprintf("0x%x", networkHeight))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		v, err := json.Marshal(a)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(string(v))
		blockHash, err := converter.Base64ToHexString(a.BlockID.Hash)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(blockHash)

		for _, txBase64 := range a.Block.Data.Txs {
			tx, err := codec.Base64ToTx(txBase64)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(tx.AuthInfo.Fee.Payer)
			fmt.Println(tx.Body.Messages[0].TypeUrl)
			txHash, err := converter.GenerateCosmosTxHashWithBase64(txBase64)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(txHash)
			if tx.Body.Messages[0].TypeUrl == "/ethermint.evm.v1.MsgEthereumTx" {
				var m evmtypes.MsgEthereumTx
				err := codec.Encoder.Unmarshal(tx.Body.Messages[0].Value, &m)
				if err != nil {
					panic(err)
				}
				fmt.Println(m.AsTransaction().Hash().Hex())

			}
		}
		os.Exit(0)
	}
}
