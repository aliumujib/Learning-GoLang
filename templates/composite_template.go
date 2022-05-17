package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func init() {
	template_ = template.Must(template.ParseGlob("templates/composite_*.html"))
}

var template_ *template.Template

func main() {
	sages := []string{"Muhd", "Ali", "Uthman", "Usman"}

	error := template_.ExecuteTemplate(os.Stdout, "composite_data_template.html", sages)

	if error != nil {
		log.Panic(error)
	}

	fmt.Println(template_)

}
