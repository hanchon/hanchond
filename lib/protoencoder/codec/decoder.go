package codec

import (
	"encoding/base64"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	protocoltypes "github.com/hanchon/hanchond/lib/types/protocol"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/evmos/v18/x/evm/types"
)

func BytesToTx(txBytes []byte) (*protocoltypes.CosmosTx, error) {
	var raw tx.TxRaw

	err := Encoder.Unmarshal(txBytes, &raw)
	if err != nil {
		return nil, err
	}

	var body tx.TxBody

	err = Encoder.Unmarshal(raw.BodyBytes, &body)
	if err != nil {
		return nil, err
	}

	var authInfo tx.AuthInfo

	err = Encoder.Unmarshal(raw.AuthInfoBytes, &authInfo)
	if err != nil {
		return nil, err
	}
	return &protocoltypes.CosmosTx{
		Body:     body,
		AuthInfo: authInfo,
	}, nil
}

func Base64ToTx(txInBase64 string) (*protocoltypes.CosmosTx, error) {
	txBytes, err := base64.StdEncoding.DecodeString(txInBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding the tx base64: %s", err.Error())
	}
	return BytesToTx(txBytes)
}

func ConvertEvmosTxToEthTx(txBase64 string) (*ethtypes.Transaction, *sdk.AccAddress, error) {
	tx, err := Base64ToTx(txBase64)
	if err != nil {
		return nil, nil, err
	}
	if len(tx.Body.Messages) == 0 {
		return nil, nil, fmt.Errorf("the transaction has no messages")
	}

	if tx.Body.Messages[0].TypeUrl != "/ethermint.evm.v1.MsgEthereumTx" {
		return nil, nil, fmt.Errorf("the message is not a eth tx")
	}

	var m evmtypes.MsgEthereumTx
	err = Encoder.Unmarshal(tx.Body.Messages[0].Value, &m)
	if err != nil {
		return nil, nil, err
	}
	signers := m.GetSigners()
	if len(signers) == 0 {
		return nil, nil, fmt.Errorf("the transaction has not signers")
	}
	from := m.GetSigners()[0]
	return m.AsTransaction(), &from, nil
}
