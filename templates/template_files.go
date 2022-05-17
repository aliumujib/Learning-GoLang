package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("templates/template.html"))
}

func main() {
	error := temp.ExecuteTemplate(os.Stdout, "template.html", 42)
	if error != nil {
		log.Fatalln(error)
	}

	fmt.Println(temp)
}
