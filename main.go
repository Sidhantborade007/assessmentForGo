// main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// Article - Our struct for all articles
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Intro   string `json:"intro"`
	Content string `json:"content"`
}
var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle Article
	json.Unmarshal(reqBody, &newArticle)

	for _, article := range Articles {
		if article.Id == id {
			article = newArticle
		}
	}

}

func requestHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	Articles = []Article{
		{Id: "1", Title: "Two states review", Intro: "Two states moview introduction", Content: "Article Content"},
		{Id: "2", Title: "Banking Finance", Intro: "Banking finance queries", Content: "Article Content2"},
	}
	requestHandler()
}