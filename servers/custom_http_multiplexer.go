package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	handleRequest(connection)
}

func handleRequest(connection net.Conn) {
	firstLine := true
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if firstLine {
			handleRoute(connection, line)
		}

		if line == "" {
			break
		}
		firstLine = false
	}
}

var Get = "GET"
var Post = "POST"

func handleRoute(connection net.Conn, line string) {
	method := strings.Fields(line)[0]
	url := strings.Fields(line)[1]

	fmt.Println("METHOD: " + method)
	fmt.Println("URL:  " + url)

	if method == Get {
		handleGet(connection, url)
	} else if method == Post {
		handlePost(connection, url)
	}
}

func handlePost(connection net.Conn, url string) {
	fmt.Println("Not yet available")
}

func handleGet(connection net.Conn, url string) {
	if url == "/" {
		serveIndex(connection)
	} else if url == "/about" {
		serveAbout(connection)
	} else if url == "/contact" {
		serveContact(connection)
	} else if url == "/apply" {
		serveApply(connection)
	}
}

func templateToString(fileUrl string, data interface{}) string {
	var doc bytes.Buffer
	template_ := template.Must(template.ParseFiles(fileUrl))
	error := template_.Execute(&doc, data)

	if error != nil {
		log.Panic(error)
	}

	result := doc.String()
	return result
}

func serveApply(connection net.Conn) {
	template_ := templateToString("servers/page_template.html", "Apply")
	fmt.Fprintf(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(template_))
	fmt.Fprintf(connection, "Content-Type: text/html\r\n")
	fmt.Fprintf(connection, "\r\n")
	fmt.Fprintf(connection, template_)
}

func serveContact(connection net.Conn) {
	template_ := templateToString("servers/page_template.html", "Contact")
	fmt.Fprintf(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(template_))
	fmt.Fprintf(connection, "Content-Type: text/html\r\n")
	fmt.Fprintf(connection, "\r\n")
	fmt.Fprintf(connection, template_)
}

func serveAbout(connection net.Conn) {
	template_ := templateToString("servers/page_template.html", "About")
	fmt.Fprintf(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(template_))
	fmt.Fprintf(connection, "Content-Type: text/html\r\n")
	fmt.Fprintf(connection, "\r\n")
	fmt.Fprintf(connection, template_)
}

func serveIndex(connection net.Conn) {
	template_ := templateToString("servers/page_template.html", "Index")
	fmt.Fprintf(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(template_))
	fmt.Fprintf(connection, "Content-Type: text/html\r\n")
	fmt.Fprintf(connection, "\r\n")
	fmt.Fprintf(connection, template_)
}
