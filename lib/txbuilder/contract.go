package txbuilder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Contract struct {
	address common.Address
	ABI     abi.ABI
}

func NewContract(address string, abi abi.ABI) Contract {
	return Contract{
		address: common.HexToAddress(address),
		ABI:     abi,
	}
}
