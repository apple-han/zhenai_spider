package main

import (
	"net/http"
	"learn/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("/Users/steven/learngo/src/learn/crawler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"/Users/steven/learngo/src/learn/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
