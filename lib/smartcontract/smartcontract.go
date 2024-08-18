package smartcontract

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func ABIPack(abiBytes []byte, method string, args ...interface{}) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to parse the ABI: %s", err.Error())
	}

	callData, err := parsedABI.Pack(method, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack the ABI: %s", err.Error())
	}

	return "0x" + hex.EncodeToString(callData), nil
}

func StringsToABIArguments(args []string) ([]interface{}, error) {
	callArgs := []interface{}{}
	for _, v := range args {
		value := strings.Split(v, ":")
		switch value[0] {
		case "a":
			callArgs = append(callArgs, common.HexToAddress(value[1]))
		default:
			return callArgs, fmt.Errorf("invalid param type")
		}
	}
	return callArgs, nil
}
