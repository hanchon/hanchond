package explorer

import (
	"context"
	"fmt"
	"sync"

	"github.com/hanchon/hanchond/playground/explorer/database"
)

type Database struct {
	queries *database.Queries
	mutex   *sync.Mutex
	ctx     context.Context
}

func NewDatabase(ctx context.Context, queries *database.Queries) *Database {
	return &Database{
		queries: queries,
		mutex:   &sync.Mutex{},
		ctx:     ctx,
	}
}

func (d *Database) GetLatestBlock() (database.Block, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.queries.GetLatestBlock(d.ctx)
}

func (d *Database) AddBlocks(blocks []Block) error {
	for _, b := range blocks {
		if err := d.AddBlock(b); err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) AddBlock(b Block) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	// TODO: move this to a sql transaction
	if _, err := d.queries.InsertBlock(d.ctx, database.InsertBlockParams{
		Height:  b.height,
		Txcount: b.txcount,
		Hash:    b.hash,
	}); err != nil {
		return fmt.Errorf("error inserting block: %s", err.Error())
	}

	for _, tx := range b.txns {
		if _, err := d.queries.InsertTransaction(d.ctx, database.InsertTransactionParams{
			Cosmoshash:  tx.cosmoshash,
			Ethhash:     tx.ethhash,
			Typeurl:     tx.typeURL,
			Sender:      tx.sender,
			Blockheight: b.height,
		}); err != nil {
			return fmt.Errorf("error inserting tx: %s", err.Error())
		}
	}

	return nil
}
