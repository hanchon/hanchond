package converter

import "strings"

// RemoveHexPrefixFromAddress returns a new string without the 0x prefix, does not modify the original string
func RemoveHexPrefixFromAddress(address string) string {
	if strings.HasPrefix(address, "0x") {
		return address[2:]
	}
	return address
}
