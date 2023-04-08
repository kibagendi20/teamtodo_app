package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//handler function which writes a byte slice containing Hello Team as the response

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello Team"))
	//log.Fatal()
}

// Add a showTodo handle function
func showTodo(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	//Fprintf formats according to a format specifier and writes to w.
	//It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(w, "Display a specific todo with ID %d...", id)

}

// Add a createTodo handle function
func createTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//this two can be replaced by:
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		// w.Header()["Date"] = nilbtw
		http.Error(w, "Method Not Allowed", 405)
		return
	}

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
