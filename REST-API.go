package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Description of second article", Content: "Article 2nd Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/all", returnAllArticles)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
	//log.Fatal(http.ListenAndServe(":8081",nil))
}

func main() {
	fmt.Prinln("Rest API v2.0 - Mux Routers")
	handleRequests()
}
