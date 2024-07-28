package converter

import "math/big"

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
