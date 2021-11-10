package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	//"path/filepath"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

type PageData struct {
	PageTitle string
	name      string
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/data", responseWithJSON)
	/*
		http.HandleFunc("/", homePage)
		http.HandleFunc("/data", responseWithJSON)
		http.HandleFunc("/testSend", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "http://localhost:8081/resource/index.html")
		})
		http.HandleFunc("/intro", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(rw, "Xin chao ban, Minh la Teo")
		})*/

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
func homePage(w http.ResponseWriter, r *http.Request) {

	//data := TodoPageData{
	//PageTitle : "My TODO list"
	//}
	//content := string("Wellcome Blockchain")\
	var content string = "Wellcome to my project " + strconv.Itoa(rand.Intn(10000))
	var_PageVariables := PageData{content, "chelsea"}

	t, err := template.ParseFiles("./resource/index.html")
	if err != nil {
		log.Println("Error: ", err)
	}
	err = t.Execute(w, var_PageVariables)
	//http.ServeFile(w,r,t)
	if err != nil {
		log.Println("Error Execute : ", err)
		log.Println("content: ", "content")
	}
}

/*func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}*/

func responseWithJSON(response http.ResponseWriter, r *http.Request) {
	result, _ := json.Marshal("Welcome to blockchain " + strconv.Itoa(rand.Intn(100000)))
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.WriteHeader(200)
	response.Write(result)
}
