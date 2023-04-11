package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/todo", app.showTodo)
	mux.HandleFunc("/todo/create", app.createTodo)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//register the file server to map with the static url
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
