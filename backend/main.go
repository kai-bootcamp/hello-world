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
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	//Xu ly bằng File server sent static resource
	fs := http.FileServer(http.Dir("resource"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))

	/*fs1 := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs1))*/
	//http.HandleFunc("/", homePage)
	// API get data là string random
	http.HandleFunc("/data", responseWithJSON)
	//Xử lý bằng cách ParseFile
	http.HandleFunc("/testRender", homePage)
	//Xử lý bằng hàm serveFile
	http.HandleFunc("/testSend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./resource/index.html")
	})
	//Get chuỗi
	http.HandleFunc("/intro", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Xin chao ban, Minh la Teo")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
func homePage(w http.ResponseWriter, r *http.Request) {

	var content string = "Wellcome to Blockchain " + strconv.Itoa(rand.Intn(10000))
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
