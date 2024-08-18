package smartcontract

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// TODO: add documentation for this function
func ABIPack(abiBytes []byte, method string, args ...interface{}) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to parse the ABI: %s\n", err.Error())
	}

	callData, err := parsedABI.Pack(method, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack the ABI: %s\n", err.Error())
	}

	return "0x" + hex.EncodeToString(callData), nil
}
