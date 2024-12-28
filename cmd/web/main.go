package main

import (
	"log"
	"net/http"
)



func main() {
	// creating a Handler mux
	mux := http.NewServeMux()

	// Controllers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Web Server
	log.Println("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux) // creating server using http.ListenAndServe on port 4000
	log.Fatal(err)
}
