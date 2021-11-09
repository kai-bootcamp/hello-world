package main

import (
	"fmt"
	"net/http"
)

const root = "/"

func main() {
	handlerWebroot()
	listenOnPort()

}

// listenOnPort
func listenOnPort() {
	http.ListenAndServe(":8080", nil)
}

// handlerWebroot tells the http package to handle all request to the web root with handler.
func handlerWebroot() {
	http.HandleFunc(root, handler)
}

// handler takes an http.ResponseWriter and http.Request as its arguments.
// http.ResponseWriter value assembles the HTTP server's response by writing to and send data to the HTTP client.
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!")
}
