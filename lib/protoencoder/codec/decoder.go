package codec

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/tx"
	protocoltypes "github.com/hanchon/vivi/lib/types/protocol"
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
