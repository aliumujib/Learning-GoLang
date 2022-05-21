package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", register)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func register(writer http.ResponseWriter, request *http.Request) {
	q := request.FormValue("q")
	io.WriteString(writer, "Your query was: "+q)
}
