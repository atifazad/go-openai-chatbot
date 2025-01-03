package main

import (
	"net/http"

	"go-chatbot/src/handlers"
)

func main() {
	http.HandleFunc("/chat", handlers.ChatHandler)
	http.Handle("/", http.FileServer(http.Dir("./src/templates")))

	http.ListenAndServe(":8080", nil)
}
