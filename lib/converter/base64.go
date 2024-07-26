package converter

import (
	"encoding/base64"
	"encoding/hex"
)

func Base64ToHexString(data string) (string, error) {
	ret, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ret), nil
}
