package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticles")
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintf(w, "Key: "+key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticles")
	// fmt.Println(r)
	// fmt.Println(r.Body)
	// fmt.Printf("%T, %T", r, r.Body)
	if r.Method != "POST" {
		json.NewEncoder(w).Encode("Get Method Not Supported")
	} else {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Fprintf(w, "%+v", string(reqBody))
		var article Article
		json.Unmarshal(reqBody, &article)
		Articles = append(Articles, article)
		json.NewEncoder(w).Encode("success")
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for idx, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:idx], Articles[idx+1:]...)
			json.NewEncoder(w).Encode("Delete Successful")
		}
	}
}

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	// add our articles route and map it to our
// 	// returnAllArticles function like so
// 	http.HandleFunc("/articles", returnAllArticles)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }

func handleRequests() {
	// creates a new instance of the nux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", returnSingleArticles).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticles).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	fmt.Println("REST API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
