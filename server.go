package main

import (
	"hello-world/wallet"
	"html/template"
	"log"
	"net/http"
	"os"
)

const root = "/"

func main() {
	handlerWebroot()
	listenOnPort()

}

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

// listenOnPort
func listenOnPort() {
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handlerWebroot tells the http package to handle all request to the web root with handlerRandomData.
func handlerWebroot() {
	http.HandleFunc(root, handlerRandomData)
}

// handlerRandomData takes an http.ResponseWriter and http.Request as its arguments.
// http.ResponseWriter value assembles the HTTP server's response by writing to and send data to the HTTP client.
func handlerRandomData(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("webpages/view.html")
	if err != nil {
		http.ServeFile(writer, request, "webpages/error_view.html")
	}
	wallet, err := wallet.RandomAddress()
	if err != nil {
		http.ServeFile(writer, request, "webpages/error_view.html")
	}
	t.Execute(writer, wallet)

}