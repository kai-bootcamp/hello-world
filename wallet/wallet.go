package wallet

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"math"
	"math/big"
)

type Wallet struct {
	Address string
	Value   string
}

var client *ethclient.Client

func init() {
	client,_ = ethclient.Dial("https://mainnet.infura.io")
}

func RandomAddress() (Wallet, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return Wallet{}, err
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)

	if err != nil {
		return Wallet{}, err
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return Wallet{}, err
	}
	address := account.Address.Hex()
	balance := checkBalanceOfAddress(address)
	return Wallet{Address: address, Value: balance.String()}, nil
}

func checkBalanceOfAddress(address string) big.Float {
	account := common.HexToAddress(address)
	balance, _ := client.BalanceAt(context.Background(), account, nil)
	ethBal := new(big.Float)
	ethBal.SetString(balance.String())
	ethValue := new(big.Float).Quo(ethBal, big.NewFloat(math.Pow10(18)))
	return *ethValue
}