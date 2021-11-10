package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
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

	//Handle home page
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	// API get data là string random
	http.HandleFunc("/data", responseWithJSON)
	//Xử lý bằng cách ParseFile
	http.HandleFunc("/testRender", homePage)
	//Xử lý bằng hàm serveFile
	http.HandleFunc("/testSend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./resource/index.html")
	})
	//func Get chuỗi
	http.HandleFunc("/intro", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Xin chao ban, Minh la Teo")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func homePage(w http.ResponseWriter, r *http.Request) {

	var content string = "Wellcome to Blockchain " + strconv.Itoa(rand.Intn(10000))
	var_PageVariables := PageData{content, "chelsea"}

	t, err := template.ParseFiles("./resource/index.html")
	if err != nil {
		log.Println("Error: ", err)
	}
	err = t.Execute(w, var_PageVariables)
	if err != nil {
		log.Println("Error Execute : ", err)
		log.Println("content: ", "content")
	}
}

func responseWithJSON(response http.ResponseWriter, r *http.Request) {
	log.Println("handle request: ", r.URL)
	var_str := "Welcome to blockchain " + strconv.Itoa(rand.Intn(100000))
	var str string = "Welcome to Blockchain\n -- " + base64.StdEncoding.EncodeToString([]byte(var_str))
	result, _ := json.Marshal(str)
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.WriteHeader(200)
	response.Write(result)
}
