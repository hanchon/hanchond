package converter

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// RemoveHexPrefixFromAddress returns a new string without the 0x prefix, does not modify the original string
func RemoveHexPrefixFromAddress(address string) string {
	if Has0xPrefix(address) {
		return address[2:]
	}
	return address
}

// Has0xPrefix from go-ethereum (copied here because it is not exported)
func Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// HexToBech32 converts from Ethereum addresses to cosmos address
func HexToBech32(address, prefix string) (string, error) {
	addr := common.HexToAddress(address)
	bech32addr, err := sdk.Bech32ifyAddressBytes(prefix, addr.Bytes())
	if err != nil {
		return "", err
	}
	return bech32addr, nil
}

// Bech32ToHex converts from cosmos address to Ethereum addresses
func Bech32ToHex(address string) (string, error) {
	bech32Prefix := strings.SplitN(address, "1", 2)[0]
	if bech32Prefix == address {
		// It was not a bech32 encoded wallet
		return "", fmt.Errorf("the wallet was not bech32 encoded")
	}
	addressBz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil {
		return "", err
	}
	return common.BytesToAddress(addressBz).Hex(), nil
}

// NormalizeAddressToHex converts from Bech32 if the address does not have the 0x prefix
func NormalizeAddressToHex(input string) (string, error) {
	if Has0xPrefix(input) {
		return input, nil
	}
	return Bech32ToHex(input)
}
