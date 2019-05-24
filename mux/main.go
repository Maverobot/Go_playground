package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is a dummy struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles is a list of Article
type Articles []Article

func postArticles(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "test post articles works")
	if err != nil {
		log.Fatal(err)
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello world!"},
		Article{Title: "Test title2", Desc: "Test description2", Content: "Hello world again!"},
	}
	fmt.Println("Endpoint hit: all articles endpoint")
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Fatal(err)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home page endpoint hit")
	if err != nil {
		log.Fatal(err)
	}
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
