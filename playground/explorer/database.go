package explorer

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/hanchon/hanchond/playground/explorer/database"
)

type Database struct {
	db      *sql.DB
	queries *database.Queries
	mutex   *sync.Mutex
	ctx     context.Context
}

func NewDatabase(ctx context.Context, db *sql.DB, queries *database.Queries) *Database {
	return &Database{
		db:      db,
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
	// SQL Transaction: to avoid corrupting the db when closing the program
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := d.queries.WithTx(tx)
	if _, err := q.InsertBlock(d.ctx, database.InsertBlockParams{
		Height:  b.height,
		Txcount: b.txcount,
		Hash:    b.hash,
	}); err != nil {
		return fmt.Errorf("error inserting block: %s", err.Error())
	}

	for _, tx := range b.txns {
		if _, err := q.InsertTransaction(d.ctx, database.InsertTransactionParams{
			Cosmoshash:  tx.cosmoshash,
			Ethhash:     tx.ethhash,
			Typeurl:     tx.typeURL,
			Sender:      tx.sender,
			Blockheight: b.height,
		}); err != nil {
			return fmt.Errorf("error inserting tx: %s", err.Error())
		}
	}

	return tx.Commit()
}
