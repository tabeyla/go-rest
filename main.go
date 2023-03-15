package main

import (
"fmt"
"log"
"net/http"
"encoding/json"
"github.com/gorilla/mux"
"io/ioutil"
)

type Article struct {
  Id string `json:"id"`
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Articles []Article



func homePage(w http.ResponseWriter, r * http.Request){

  fmt.Fprintf(w, "Welcome to the homepage")
  fmt.Println("Endpoint hit: homePage")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
 fmt.Println("Endpoint hit: createNewArticle")
 var article Article
 
 reqBody, _ := ioutil.ReadAll(r.Body)
 json.Unmarshal(reqBody, &article)

 Articles = append(Articles, article)

 json.NewEncoder(w).Encode(article)
}
func returnSingleArticle(w http.ResponseWriter, r * http.Request) {
 vars := mux.Vars(r)
 key := vars["id"]

 for _,article := range Articles {
  if article.Id == key {
  json.NewEncoder(w).Encode(article)
 }
}
 fmt.Println("Invalid id "+ key)
}

func handleRequestsMux() {

 router := mux.NewRouter().StrictSlash(true)
 router.HandleFunc("/", homePage)
 router.HandleFunc("/articles", returnAllArticles)
 router.HandleFunc("/article/{id}", returnSingleArticle)
 router.HandleFunc("/article", createNewArticle).Methods("POST")
 log.Fatal(http.ListenAndServe(":10000", router))
}

func handleRequests() {

  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
 fmt.Println("Endpoint hit: returnAllArticles")
 json.NewEncoder(w).Encode(Articles)
}

func main() {
// handleRequests()

 Articles = []Article{
	Article{"1", "Hello", "Hello Description", "Hello Content"},
	Article{"2", "Hello2", "Hello2 Description", "Hello2 Content"},
   }

 handleRequestsMux()
}
