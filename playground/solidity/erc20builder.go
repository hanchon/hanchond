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

func GenerateWrappedCoinContract(name, symbol, decimals string) string {
	name = StringToTitle(name)
	symbol = strings.ToUpper(symbol)
	return fmt.Sprintf(`pragma solidity ^0.4.18;
contract %s {
    string public name = "%s";
    string public symbol = "%s";
    uint8 public decimals = %s;

    event Approval(address indexed src, address indexed guy, uint wad);
    event Transfer(address indexed src, address indexed dst, uint wad);
    event Deposit(address indexed dst, uint wad);
    event Withdrawal(address indexed src, uint wad);

    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    function() public payable {
        deposit();
    }

    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        Deposit(msg.sender, msg.value);
    }

    function withdraw(uint wad) public {
        require(balanceOf[msg.sender] >= wad);
        balanceOf[msg.sender] -= wad;
        msg.sender.transfer(wad);
        Withdrawal(msg.sender, wad);
    }

    function totalSupply() public view returns (uint) {
        return this.balance;
    }

    function approve(address guy, uint wad) public returns (bool) {
        allowance[msg.sender][guy] = wad;
        Approval(msg.sender, guy, wad);
        return true;
    }

    function transfer(address dst, uint wad) public returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    function transferFrom(
        address src,
        address dst,
        uint wad
    ) public returns (bool) {
        require(balanceOf[src] >= wad);

        if (src != msg.sender && allowance[src][msg.sender] != uint(-1)) {
            require(allowance[src][msg.sender] >= wad);
            allowance[src][msg.sender] -= wad;
        }

        balanceOf[src] -= wad;
        balanceOf[dst] += wad;

        Transfer(src, dst, wad);

        return true;
    }
}`, name, name, symbol, decimals)
}
