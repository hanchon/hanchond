package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/txbuilder"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"golang.org/x/exp/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage [mnemonic] [web3_endpoint] [home_dir]")
		os.Exit(1)
	}
	mnemonic := os.Args[1]
	web3Endpoint := os.Args[2]
	homeDir := os.Args[3]
	// This is needed because it will build the erc20 contract with the solc version downloaded with `build-solc`
	filesmanager.SetBaseDir(homeDir)

	valWallet := txbuilder.NewSimpleWeb3WalletFromMnemonic(mnemonic, web3Endpoint)
	client := requester.NewClient().WithUnsecureWeb3Endpoint(web3Endpoint)

	erc20sAddress := []string{}

	// Create some erc20s. Every deployment will wait until the tx is included in a block to ensure the correct deployment of the contract
	for range 2 {
		initialAmount := "1000000"
		txHash, err := solidity.BuildAndDeployERC20Contract(randString(7), randString(3), initialAmount, false, valWallet.TxBuilder, 1_000_000)
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
	wallets := []*txbuilder.SimpleWeb3Wallet{}
	for range totalWallets {
		w := txbuilder.NewSimpleWeb3Wallet(web3Endpoint)
		wallets = append(wallets, w)

		if _, err := valWallet.TxBuilder.SendCoins(
			w.Address.Hex(),
			big.NewInt(1_000_000_000_000_000_000),
		); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, erc20Wallet := range erc20sAddress {
			callData, err := solidity.ERC20TransferCallData(w.Address.Hex(), "100000")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			to := common.HexToAddress(erc20Wallet)
			if _, err := valWallet.TxBuilder.SendTx(valWallet.Address, &to, big.NewInt(0), 200_000, callData, valWallet.PrivKey); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}

	for {
		// TODO: add height counter
		txHash, _ := sendRandomTransaction(erc20sAddress, wallets)
		fmt.Println(txHash)
	}
}

func sendRandomTransaction(erc20Address []string, wallets []*txbuilder.SimpleWeb3Wallet) (string, error) {
	from := wallets[rand.Intn(len(wallets))]
	toWallet := wallets[rand.Intn(len(wallets))]
	erc20 := erc20Address[rand.Intn(len(erc20Address))]

	callData, err := solidity.ERC20TransferCallData(toWallet.Address.Hex(), "1")
	if err != nil {
		return "", err
	}

	to := common.HexToAddress(erc20)
	return from.TxBuilder.SendTx(from.Address, &to, big.NewInt(0), 200_000, callData, from.PrivKey)
}
