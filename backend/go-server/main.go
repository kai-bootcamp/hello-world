package main

// 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Global = "myvalue"

type BlockChainPage struct {
	Title   string
	Address string
	Success bool
}

var client, err = ethclient.Dial("https://mainnet.infura.io/v3/f65c68ec7a2a4660aec71b82727181f6")

const PORT = "8080"
const ERROR404 = "404 page not found !!!"

func getBalance(address string) string {
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(balance) // 25893180161173005034
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	return ethValue.String()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello/" {
		http.Error(w, "wdwdwdw", http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("./static/basictemplating.html")
	if r.Method != http.MethodPost {
		p := BlockChainPage{Title: "Test Page"}
		t.Execute(w, p)
		return
	}
	p := BlockChainPage{Title: "Thansk for your input", Address: getBalance(r.FormValue("address")), Success: true}
	t.Execute(w, p)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/blockChain/", func(w http.ResponseWriter, r *http.Request) {

		t, _ := template.ParseFiles("./static/basictemplating.html")
		if r.Method != http.MethodPost {
			//p := BlockChainPage{Title: "Test Page"}
			t.Execute(w, nil)
			return
		}
		p := BlockChainPage{Title: "Test", Address: r.FormValue("address")}
		t.Execute(w, struct{ Success bool }{true})
		log.Print(p.Address)
		//t.Execute(w, p)

	})

	fmt.Print("Starting server at port: ", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
