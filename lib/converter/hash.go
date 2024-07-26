package converter

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func GenerateCosmosTxHash(txBytes []byte) string {
	hash := sha256.Sum256(txBytes)
	return hex.EncodeToString(hash[:])
}

func GenerateCosmosTxHashWithBase64(txInBase64 string) (string, error) {
	txBytes, err := base64.StdEncoding.DecodeString(txInBase64)
	if err != nil {
		return "", err
	}
	return GenerateCosmosTxHash([]byte(txBytes)), nil
}
