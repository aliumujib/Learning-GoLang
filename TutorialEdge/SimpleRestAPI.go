package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func saveArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Save articles endpoint hit")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Title of first article", Desc: "Description of first article", Content: "Content of first article"},
	}
	fmt.Println("This endpoint returns the articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage endpoint hit")
}

func handleRequests() {
	fmt.Println("Up and running")
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/articles", allArticles).Methods("GET")
	muxRouter.HandleFunc("/saveArticles", saveArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8003", muxRouter))
}

func main() {
	handleRequests()
}
