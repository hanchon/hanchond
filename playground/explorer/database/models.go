// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type Block struct {
	ID         interface{}
	Height     int64
	Txcount    int64
	Hash       string
	Parenthash string
}

type Transaction struct {
	ID          interface{}
	Cosmoshash  string
	Ethhash     string
	Content     string
	Sender      string
	Blockheight int64
}
