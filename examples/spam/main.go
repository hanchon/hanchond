package main

import (
	"crypto/ecdsa"
	_ "embed"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/smartcontract"
	"github.com/hanchon/hanchond/lib/txbuilder"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"golang.org/x/exp/rand"
)

type Wallet struct {
	mnemonic  string
	address   common.Address
	privKey   *ecdsa.PrivateKey
	txbuilder *txbuilder.TxBuilder
}

func NewWallet(mnemonic string, web3Endpoint string) *Wallet {
	w, a, err := txbuilder.WalletFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println("error generating wallet 1", err.Error())
		os.Exit(1)
	}
	privKey, err := w.PrivateKey(a)
	if err != nil {
		fmt.Println("error generating wallet 1", err.Error())
		os.Exit(1)
	}

	return &Wallet{
		mnemonic:  mnemonic,
		address:   a.Address,
		privKey:   privKey,
		txbuilder: txbuilder.NewSimpleTxBuilder(mnemonic, web3Endpoint),
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	filesmanager.SetBaseDir("/Users/hanchon/.hanchond")
	// if len(os.Args) < 3 {
	// 	fmt.Println("usage [mnemonic] [web3_endpoint]")
	// 	os.Exit(1)
	// }
	// mnemonic := os.Args[1]
	// web3Endpoint := os.Args[2]
	mnemonic := "marble plug replace plastic close crawl mandate bundle flat verb tortoise vessel elbow clever harsh roof crawl diary skin veteran salmon hat oxygen noise"
	web3Endpoint := "http://localhost:49901"

	valWallet := NewWallet(mnemonic, web3Endpoint)
	client := requester.NewClient().WithUnsecureWeb3Endpoint(web3Endpoint)

	erc20sAddress := []string{}

	// Create some erc20s. Every deployment will wait until the tx is included in a block to ensure the correct deployment of the contract
	for range 2 {
		initialAmount := "1000000"
		txHash, err := solidity.BuildDeployERC20Contract(randString(7), randString(3), initialAmount, false, valWallet.txbuilder, 1_000_000)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		address, err := client.GetContractAddress(txHash)
		if err == nil {
			erc20sAddress = append(erc20sAddress, address)
		}
		fmt.Println("contract deployed:", address)
	}

	// Create the wallets and send coins
	totalWallets := 100
	wallets := []*Wallet{}
	for range totalWallets {
		mnemonic, _ := txbuilder.NewMnemonic()
		w := NewWallet(mnemonic, web3Endpoint)
		wallets = append(wallets, w)

		// TODO: add simple coin transfer method
		v, err := valWallet.txbuilder.SendTx(
			valWallet.address,
			&w.address,
			big.NewInt(1_000_000_000_000_000_000),
			25_000,
			[]byte{},
			valWallet.privKey,
		)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("address", w.address.Hex())
		fmt.Println("erc20", erc20sAddress[0])

		for _, erc20Wallet := range erc20sAddress {
			callData, err := ERC20TransferCallData(w.address.Hex(), "100000")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			to := common.HexToAddress(erc20Wallet)
			v, err = valWallet.txbuilder.SendTx(valWallet.address, &to, big.NewInt(0), 200_000, callData, valWallet.privKey)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Println(v)
		}
	}

	for {
		txHash, _ := sendRandomTransaction(erc20sAddress, wallets)
		fmt.Println(txHash)
	}
}

func sendRandomTransaction(erc20Address []string, wallets []*Wallet) (string, error) {
	fmt.Println("send")
	from := wallets[rand.Intn(len(wallets))]
	toWallet := wallets[rand.Intn(len(wallets))]
	erc20 := erc20Address[rand.Intn(len(erc20Address))]

	callData, err := ERC20TransferCallData(toWallet.address.Hex(), "1")
	if err != nil {
		return "", err
	}

	to := common.HexToAddress(erc20)
	return from.txbuilder.SendTx(from.address, &to, big.NewInt(0), 200_000, callData, from.privKey)
}

var erc20abi = `[{
        "constant": false,
        "inputs": [
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "transfer",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }]`

func ERC20TransferCallData(address string, amount string) ([]byte, error) {
	params := []string{"a:" + address, "n:" + amount}
	callArgs, err := smartcontract.StringsToABIArguments(params)
	if err != nil {
		return []byte{}, err
	}

	return smartcontract.ABIPackRaw([]byte(erc20abi), "transfer", callArgs...)
}
