package explorer

import (
	"context"
	"fmt"
	"log"

	"github.com/hanchon/hanchond/playground/explorer/database"
)

type ExplorerClient struct {
	web3Endpoint   string
	cosmosEndpoint string

	ctx context.Context

	queries *database.Queries
}

func NewLocalExplorerClient(web3Port, cosmosPort int, dataPath string) *ExplorerClient {
	queries, err := database.InitExplorerDatabase(context.Background(), dataPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ExplorerClient{
		web3Endpoint:   fmt.Sprintf("http://localhost:%d", web3Port),
		cosmosEndpoint: fmt.Sprintf("http://localhost:%d", cosmosPort),
		queries:        queries,
		ctx:            context.Background(),
	}
}

func (e *ExplorerClient) ProcessBlocks() {

}
