package solidity

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hanchon/hanchond/lib/smartcontract"
	"github.com/hanchon/hanchond/lib/txbuilder"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const erc20TransferABI = `[{ "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_value", "type": "uint256" } ], "name": "transfer", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }]`

func ERC20TransferCallData(address string, amount string) ([]byte, error) {
	params := []string{"a:" + address, "n:" + amount}
	callArgs, err := smartcontract.StringsToABIArguments(params)
	if err != nil {
		return []byte{}, err
	}
	return smartcontract.ABIPackRaw([]byte(erc20TransferABI), "transfer", callArgs...)
}

// BuildAndDeployERC20Contract will save the temp usings using the filesmanager. Init the home folder before running the function
func BuildAndDeployERC20Contract(name, symbol, initialAmount string, isWrapped bool, builder *txbuilder.TxBuilder, gasLimit uint64) (string, error) {
	// Clone openzeppelin if needed
	path, err := DownloadDep("https://github.com/OpenZeppelin/openzeppelin-contracts", "v5.0.2", "openzeppelin")
	if err != nil {
		return "", err
	}

	// Set up temp folder
	if err := filesmanager.CleanUpTempFolder(); err != nil {
		return "", fmt.Errorf("could not clean up the temp folder:%s", err.Error())
	}

	folderName := "erc20builder"
	if err := filesmanager.CreateTempFolder(folderName); err != nil {
		return "", fmt.Errorf("could not create the temp folder:%s", err.Error())
	}

	contract := ""
	solcVersion := "0.8.25"
	switch isWrapped {
	case false:
		// Normal ERC20
		contract = GenerateERC20Contract(path, name, symbol, initialAmount)
	case true:
		// Wrapping base denom, use WETH9
		contract = GenerateWrappedCoinContract(name, symbol, "18")
		solcVersion = "0.4.18"
	}

	contractPath := filesmanager.GetBranchFolder(folderName) + "/mycontract.sol"
	if err := filesmanager.SaveFile([]byte(contract), contractPath); err != nil {
		return "", fmt.Errorf("could not save the contract file:%s", err.Error())
	}

	// Compile the contract
	err = CompileWithSolc(solcVersion, contractPath, filesmanager.GetBranchFolder(folderName))
	if err != nil {
		return "", fmt.Errorf("could not compile the erc20 contract:%s", err.Error())
	}

	bytecode, err := filesmanager.ReadFile(filesmanager.GetBranchFolder(folderName) + "/" + StringToTitle(name) + ".bin")
	if err != nil {
		return "", fmt.Errorf("error reading the bytecode file:%s", err.Error())
	}

	bytecode, err = hex.DecodeString(string(bytecode))
	if err != nil {
		return "", fmt.Errorf("error converting bytecode to []byte:%s", err.Error())
	}

	txHash, err := builder.DeployContract(0, bytecode, gasLimit)
	if err != nil {
		return "", fmt.Errorf("error sending the transaction:%s", err.Error())
	}
	return txHash, nil
}

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
