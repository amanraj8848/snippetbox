package main

import(
    "net/http"
    "strconv"
    "fmt"
	"log"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	// parse the home.page.tmpl file using template.ParseFiles
	// this will allow us to execute the template and fill in the dynamic data
	// if there is an error parsing the file, log the error and return a 500 status code
	ts,err := template.ParseFiles("ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Execute the template and write the output to the http.ResponseWriter
	// Passing 'nil' as the data since there is no dynamic data to inject
	err = ts.Execute(w, nil)
	if err != nil {
		// Log the error message if template execution fails
		log.Println(err.Error())
		// Send a 500 Internal Server Error response to the client
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Write([]byte("hello world"))
	
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	// lets get the value of the id parameter in the url (eg. /snippet?id=1)
	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
	w.Write([]byte("snippet"))
}



func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		w.Header().Set("Allow", "POST") // adds a header named "Allow" with the value "POST"
		// w.WriteHeader(405)  // can only call WriteHeader once per respsone, must call w.WriteHeader() before any call to w.Write()
		// w.Write([]byte("method not allowed")) // If you don’t call w.WriteHeader() explicitly, then the first call to w.Write() will automatically send a 200 OK status code to the user
		http.Error(w, "method not allowed", 405)  // this is the same as calling w.WriteHeader(405) with w.Write([]byte("method not allowed"))
		return 
	}
	
	// w.Header().Set("Content-Type", "application/json")  // used to set the response type to JSON w.Write([]byte(`{"name":"Alex"}`))
	w.Write([]byte("creating a new snippet"))
}
