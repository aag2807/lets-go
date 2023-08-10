package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("src/ui/static"))
	handler := Handler{}
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("/", http.HandlerFunc(handler.HomeHandler))
	mux.Handle("/snippet/view", http.HandlerFunc(handler.ViewSnippetHandler))
	mux.Handle("/snippet/create", http.HandlerFunc(handler.CreateSnippetHanlder))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
