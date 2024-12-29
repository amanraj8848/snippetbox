package main

import (
	"log"
	"net/http"
)



func main() {
	// creating a Handler mux
	mux := http.NewServeMux()

/*************  ✨ Codeium Command 🌟  *************/
	// Create a file server handler to serve static files from the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
/******  52c2c6ae-95ec-456e-8232-e60e6e7d2f67  *******/

	// Controllers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	// Use the mux.Handle() function to register the file server as the handler 
// all URL paths that start with "/static/". For matching paths, we strip the
 // "/static" prefix before the request reaches the file server. 
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Web Server
	log.Println("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux) // creating server using http.ListenAndServe on port 4000
	log.Fatal(err)
}
