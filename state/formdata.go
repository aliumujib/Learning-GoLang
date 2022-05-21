package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	FirstName  string
	LastName   string
	Subscribed string
}

func main() {
	http.HandleFunc("/register", search)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func search(writer http.ResponseWriter, request *http.Request) {
	person := Person{
		request.FormValue("first"),
		request.FormValue("last"),
		request.FormValue("subscribe"),
	}

	tpl := template.Must(template.ParseGlob("state/templates/*"))

	error := tpl.ExecuteTemplate(writer, "index.html", person)
	if error != nil {
		http.Error(writer, error.Error(), 500)
	}
}
