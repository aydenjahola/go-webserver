package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello handler function
func helloHandler(w http.ResponseWriter, r *http.Request) { // w is the response writer and r is the request
	if r.URL.Path != "/hello" { // if the URL path is not /hello
		http.Error(w, "404 not found.", http.StatusNotFound) // return 404 not found, StatusNotFound is a constant in the http package
		return
	}

	if r.Method != "GET" { // check if the request method is not GET
		http.Error(w, "Method is not supported.", http.StatusNotFound) // return Method is not supported so we dont allow other methods to be used on this route
		return
	}

	fmt.Fprint(w, "Hello!") // if the request is successful, we print Hello! to the response writer w.

}

// Form handler function
func formHandler(w http.ResponseWriter, r *http.Request) { // w is the response writer and r is the request
	if err := r.ParseForm(); err != nil { // ParseForm parses the raw query from the URL and updates r.Form. For POST, PUT, and PATCH requests, it also parses the request body as a form and puts the results into both r.PostForm and r.Form.
		fmt.Fprintf(w, "ParseForm() err: %v", err) // if there is an error while parsing the form, we print the error message to the response writer w. %v is a placeholder for the error message
		return
	}
	
	// print the request method to the console
	fmt.Fprintf(w, "POST request successful\n") // if the request is successful, we print POST request successful to the response writer w.

	// Form is a struct that represents a parsed form. It is not a map of key/value pairs, but a map of key to a list of values for that key.
	// ParseForm populates r.Form and r.PostForm.
	// For all requests, ParseForm parses the raw query from the URL and updates r.Form.
	// For POST, PUT, and PATCH requests, it also parses the request body as a form and puts the results into both r.PostForm and r.Form.
	r.ParseForm()

	// fmt.Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(w, "Hello %s!", r.FormValue("name")) // if the request is successful, we print Hello %s! to the response writer w. %s is a placeholder for the name value
}

// Main function is the entry point of the program
func main() {
	// short form assingment operator := -> declares and assigns a value to a variable
	fileServer := http.FileServer(http.Dir("./static")) // this will serve the files in the static directory
	http.Handle("/", fileServer) // this will serve the files in the static directory

	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/form", formHandler) // this will serve the files in the static directory

	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/hello", helloHandler) // this will serve the files in the static directory

	// fmt.Printf is used to print the message to the console
	fmt.Printf("Starting server at port 8080\n")

	// if we encounter an error while starting the server, we log the error message
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}