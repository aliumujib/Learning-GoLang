package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is a dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	var doc bytes.Buffer
	temp := template.Must(template.ParseFiles("templates/template.html"))
	temp.Execute(&doc, "Aliu Abdul-Mujeeb")
	io.WriteString(res, doc.String())
}

func main() {
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
