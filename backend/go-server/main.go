package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"
const ERROR404 = "404 page not found !!!"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "wdwdwdw", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not Support", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello BlockChain!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port: ", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
