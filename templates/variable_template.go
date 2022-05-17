package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	sex := "Male"

	template_ := template.Must(template.ParseGlob("templates/*.html"))

	error := template_.ExecuteTemplate(os.Stdout, "variable_template.html", sex)

	if error != nil {
		log.Panic(error)
	}

	fmt.Println(template_)

}
