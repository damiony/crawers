package main

import (
	"frontend/controller"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./view")))

	http.Handle("/search", controller.CreateSearchResultHandler("./view/template.html"))

	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		log.Fatal(err)
	}
}
