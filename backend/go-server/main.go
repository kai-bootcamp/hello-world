package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

func helloHandler(res http.ResponseWriter, req *http.Request)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port: ", PORT)
	if err := http.ListenAndServe(":8080", nil); err != nill {
		panic(err)
	}

}
