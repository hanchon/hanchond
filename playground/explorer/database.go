package explorer

import (
	"context"
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
