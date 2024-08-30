-- name: GetLatestBlock :one
SELECT * FROM blocks ORDER BY height DESC LIMIT 1;

-- name: DeleteBlockByID :exec
DELETE FROM blocks WHERE id = ?;

-- name: InsertBlock :one
INSERT INTO blocks(
    height,  txcount,  hash
) VALUES (
    ?, ?, ?
)
RETURNING id;

-- name: InsertTransaction :one
INSERT INTO transactions(
    cosmoshash, ethhash, typeurl, sender, blockheight
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING id;

-- name: GetTransactions :many
SELECT * FROM transactions;

-- name: GetLimitedTransactions :many
SELECT * FROM transactions ORDER BY id DESC LIMIT ?;

