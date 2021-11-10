package wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wemeetagain/go-hdwallet"
	"io/ioutil"
	"math"
	"net/http"
)

type Wallet struct {
	Address    string
	Balance    string
	PrivateKey string
}

func RandomAddress() (Wallet, error) {
	seed, err := hdwallet.GenSeed(128)
	if err != nil {
		return Wallet{}, err
	}
	masterKey := hdwallet.MasterKey(seed)
	childKey, _ := masterKey.Child(0)
	childPub := childKey.Pub()
	address := childPub.Address()
	balance := checkBalanceOfAddress(address)
	keyString := hex.EncodeToString(childKey.Key)[2:]
	return Wallet{Address: address, PrivateKey: keyString, Balance: fmt.Sprintf("%g", balance)}, nil
}

func checkBalanceOfAddress(address string) float64 {
	response, err := http.Get("https://blockstream.info/api/address/" + address)
	if err == nil {
		if response.StatusCode == 200 {
			result := response.Body
			defer result.Close()
			data, err := ioutil.ReadAll(result)
			if err != nil {
				return 0.0
			}
			mapData := make(map[string]interface{})
			json.Unmarshal(data, &mapData)
			chainStatData := mapData["chain_stats"].(map[string]interface{})
			mempoolStats := mapData["mempool_stats"].(map[string]interface{})
			fundedSum := chainStatData["funded_txo_sum"].(float64) + mempoolStats["funded_txo_sum"].(float64)
			spentSum := chainStatData["spent_txo_sum"].(float64) + mempoolStats["spent_txo_sum"].(float64)
			return (fundedSum - spentSum) / math.Pow10(8)
		}
	}
	return 0.0
}
