package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello Wolrd %q", html.EscapeString(r.URL.Path))
		http.ServeFile(w, r, "./index.html")
		//http.FileServer(http.Dir("./"))
	})

	http.HandleFunc("/intro", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Xin chao ban, Minh la Teo")
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, r.URL.Path[1:])
		http.ServeFile(w, r, "./index.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
