package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article //array of article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")

	//NewEncoder(w) -> w'ye yazan bir encoder döndürür.
	//Encode() -> JSON kodlamasını yazar.
	//list of articles
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post Endpoint Worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	//HandleFunc enable to multiple Request Routing. (route,handle function)
	//http.HandleFunc("/", homePage)

	myRouter := mux.NewRouter().StrictSlash(true)

	//gorilla mux kullanarak, path'e göre fonksiyon çalıştırabiliriz ve,
	//spesifik olarak hangi fonksiyonu kullanarak, belirtilen path'e
	//ne isteği(GET-POST-PUT) yapacağımızı  söyleyebiliriz.

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")

	// log.Fatal() -> print specified message with timestamp on the console screen.
	//http.ListenAndServe -> opens the server port, and blocks forever waiting for clients.
	//If it fails to open the port, the log.Fatal call will report the problem and exit the program.
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
