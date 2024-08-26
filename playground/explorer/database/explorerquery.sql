-- name: GetLatestBlock :one
SELECT * FROM blocks ORDER BY height DESC LIMIT 1;

-- name: DeleteBlockByID :exec
DELETE FROM blocks WHERE id = ?;

-- name: InsertBlock :one
INSERT INTO blocks(
    height, time, txcount, totalValue, proposer, gasused, gaslimit, basefee, hash, parenthash
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id;
