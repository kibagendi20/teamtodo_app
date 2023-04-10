package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//handler function which writes a byte slice containing Hello Team as the response

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//initializing slice containing the paths to the two files
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// w.Write([]byte("Hello Team"))
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
