package explorer

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/hanchon/hanchond/playground/explorer/database"
)

type cache struct {
	valid  bool
	blocks []database.Block
	txns   []database.Transaction
}

type Database struct {
	db      *sql.DB
	queries *database.Queries
	mutex   *sync.Mutex
	ctx     context.Context

	cache cache
}

func NewDatabase(ctx context.Context, db *sql.DB, queries *database.Queries) *Database {
	return &Database{
		db:      db,
		queries: queries,
		mutex:   &sync.Mutex{},
		ctx:     ctx,

		cache: cache{valid: false},
	}
}

func (d *Database) GetDisplayInfo(limit int) ([]database.Block, []database.Transaction, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.cache.valid {
		return d.cache.blocks, d.cache.txns, nil
	}

	blocks, err := d.queries.GetLimitedBlocks(d.ctx, int64(limit))
	if err != nil {
		return []database.Block{}, []database.Transaction{}, err
	}
	txns, err := d.queries.GetLimitedTransactions(d.ctx, int64(limit))
	if err != nil {
		return []database.Block{}, []database.Transaction{}, err
	}

	d.cache.valid = false
	d.cache.blocks = blocks
	d.cache.txns = txns

	return blocks, txns, nil
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
	d.cache.valid = false
	// SQL Transaction: to avoid corrupting the db when closing the program
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	//nolint: errcheck
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
