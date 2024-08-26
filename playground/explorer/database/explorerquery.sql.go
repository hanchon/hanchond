// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: explorerquery.sql

package database

import (
	"context"
)

const deleteBlockByID = `-- name: DeleteBlockByID :exec
DELETE FROM blocks WHERE id = ?
`

func (q *Queries) DeleteBlockByID(ctx context.Context, id interface{}) error {
	_, err := q.db.ExecContext(ctx, deleteBlockByID, id)
	return err
}

const getLatestBlock = `-- name: GetLatestBlock :one
SELECT id, height, time, txcount, totalvalue, proposer, gasused, gaslimit, basefee, hash, parenthash FROM blocks ORDER BY height DESC LIMIT 1
`

func (q *Queries) GetLatestBlock(ctx context.Context) (Block, error) {
	row := q.db.QueryRowContext(ctx, getLatestBlock)
	var i Block
	err := row.Scan(
		&i.ID,
		&i.Height,
		&i.Time,
		&i.Txcount,
		&i.Totalvalue,
		&i.Proposer,
		&i.Gasused,
		&i.Gaslimit,
		&i.Basefee,
		&i.Hash,
		&i.Parenthash,
	)
	return i, err
}

const insertBlock = `-- name: InsertBlock :one
INSERT INTO blocks(
    height, time, txcount, totalValue, proposer, gasused, gaslimit, basefee, hash, parenthash
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id
`

type InsertBlockParams struct {
	Height     int64
	Time       interface{}
	Txcount    int64
	Totalvalue interface{}
	Proposer   string
	Gasused    interface{}
	Gaslimit   interface{}
	Basefee    interface{}
	Hash       string
	Parenthash string
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, insertBlock,
		arg.Height,
		arg.Time,
		arg.Txcount,
		arg.Totalvalue,
		arg.Proposer,
		arg.Gasused,
		arg.Gaslimit,
		arg.Basefee,
		arg.Hash,
		arg.Parenthash,
	)
	var id interface{}
	err := row.Scan(&id)
	return id, err
}
