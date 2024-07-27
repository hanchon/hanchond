package converter

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
