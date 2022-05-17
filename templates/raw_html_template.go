package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Aliu Abdul-Mujeeb"

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title> Hello world! </title>
	</head>

	<body>
	<h1> ` + name + ` </h1>
	</body>
	</html>
	`

	nf, err := os.Create("templates/index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(template))
}
