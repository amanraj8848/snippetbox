package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello world"))
	
}

func snippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("snippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		w.Header().Set("Allow", "POST") // tells the user which methods are allowed, used along with w.WriteHeader(405)
		// w.WriteHeader(405)  // can only call WriteHeader once per respsone, must call w.WriteHeader() before any call to w.Write()
		// w.Write([]byte("method not allowed")) // If you don’t call w.WriteHeader() explicitly, then the first call to w.Write() will automatically send a 200 OK status code to the user
		http.Error(w, "method not allowed", 405)  // this is the same as calling w.WriteHeader(405) with w.Write([]byte("method not allowed"))
		return 
	}


	w.Write([]byte("creating a new snippet"))
}

func main() {
	// creating a Handler mux
	mux := http.NewServeMux()

	// Controllers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Web Server
	log.Println("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux) // creating server using http.ListenAndServe on port 4000
	log.Fatal(err)
}
