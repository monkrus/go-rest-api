package main

import (
	"fmt"
	"log"
	"net/http"
)

/*type Article struct {
Title string `json:"Title"`
Desc string `json:"desc"`
Content string `json:"content"`
}*/

/*type Articles[] Article {
	Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}*/

/*func allArticles( w http.ResponseWriter, r *http.Request) {
fmt.Println("Endpoint Hit: All Articles Endpoint")
json.NewCoder(w).Encode(articles)
}*/

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	//fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
