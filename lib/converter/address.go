package converter

import (
	"math/big"
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

// HexStringToDecimal assumes that value is a valid hex string
func HexStringToDecimal(value string) string {
	bigValue := new(big.Int)
	if Has0xPrefix(value) {
		value = value[2:]
	}
	_, ok := bigValue.SetString(value, 16)
	if !ok {
		return ""
	}
	return bigValue.Text(10)
}

// DecimalStringToHex assumes that value is a valid decimal string
func DecimalStringToHex(value string) string {
	bigValue := new(big.Int)
	_, ok := bigValue.SetString(value, 10)
	if !ok {
		return ""
	}
	return "0x" + bigValue.Text(16)
}
