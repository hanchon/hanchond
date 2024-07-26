package protocol

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type CosmosTx struct {
	Body     tx.TxBody
	AuthInfo tx.AuthInfo
}

type MsgForDB struct {
	ID    int64
	Table string
}
