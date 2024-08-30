package converter

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/hanchon/hanchond/lib/protoencoder/codec"
	"golang.org/x/crypto/sha3"
)

func GenerateCosmosTxHash(txBytes []byte) string {
	hash := sha256.Sum256(txBytes)
	return hex.EncodeToString(hash[:])
}

func GenerateEthTxHash(txBytes []byte) (string, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write(txBytes)
	if err != nil {
		return "", err
	}
	buf := hash.Sum(nil)
	return hex.EncodeToString(buf), nil
}

func GenerateEthTxHashFromEvmosTx(txBase64 string) (string, error) {
	tx, _, err := codec.ConvertEvmosTxToEthTx(txBase64)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func GenerateCosmosTxHashWithBase64(txInBase64 string) (string, error) {
	txBytes, err := base64.StdEncoding.DecodeString(txInBase64)
	if err != nil {
		return "", err
	}
	return GenerateCosmosTxHash(txBytes), nil
}
