package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type PersonWithWriteUp struct {
	FirstName      string
	LastName       string
	Subscribed     string
	WriteUpContent string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("state/templates/*"))
}

func main() {
	http.HandleFunc("/upload", readFile)
	http.ListenAndServe(":8080", nil)
}

func readFile(writer http.ResponseWriter, req *http.Request) {
	var fileContent string
	if req.Method == http.MethodPost {
		file, header, error := req.FormFile("q")

		if error != nil {
			http.Error(writer, error.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		fmt.Println("\nfile:", file, "\nheader: ", header, "\nerror:", error)

		byteStream, error := ioutil.ReadAll(file)
		if error != nil {
			http.Error(writer, error.Error(), http.StatusInternalServerError)
			return
		}

		fileContent = string(byteStream)
		fmt.Println("Content: ", fileContent)
	}

	personWithWriteUp := PersonWithWriteUp{
		req.FormValue("first"),
		req.FormValue("last"),
		req.FormValue("subscribe"),
		fileContent,
	}

	//writer.Header().Set("Content-Type", "text/html: charset=utf-8")
	error := tpl.ExecuteTemplate(writer, "index-file.html", personWithWriteUp)

	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
	}
}
