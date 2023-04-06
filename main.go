package main

import (
	"log"
	"net/http"
)

//handler function which writes a byte slice containing Hello Team as the response

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Team"))
	//log.Fatal()
}

// Add a showTodo handle function
func showTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific todo"))
}

// Add a createTodo handle function
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creates a todo"))
}

// Use the http.NewServeMux() function to initialize a new servemux, then
// register the home function as the handler for the "/" URL pattern.
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/todo", showTodo)
	mux.HandleFunc("/todo/create", createTodo)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
