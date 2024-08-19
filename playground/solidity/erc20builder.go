package solidity

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func StringToTitle(in string) string {
	return cases.Title(language.English, cases.Compact).String(in)
}

func GenerateERC20Contract(openzeppelinPath, name, symbol, initialAmount string) string {
	name = StringToTitle(name)
	symbol = strings.ToUpper(symbol)
	return fmt.Sprintf(`// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.20;

import "%s/contracts/token/ERC20/ERC20.sol";
import "%s/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract %s is ERC20, ERC20Permit {
    constructor() ERC20("%s", "%s") ERC20Permit("%s") {
        _mint(msg.sender, %s * 10 ** decimals());
    }
}
`, openzeppelinPath, openzeppelinPath, name, name, symbol, name, initialAmount)
}
