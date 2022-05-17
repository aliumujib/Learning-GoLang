package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type NewsIndex struct {
	NewsItems []News `xml:"url"`
}

type News struct {
	Urls         string `xml:"loc"`
	LastModDates string `xml:"lastmod"`
	Titles       string `xml:"news>title"`
	PublishDates string `xml:"news>publication_date"`
}

type NewsArticle struct {
	Url          string
	LastModDate  string
	Titles       string
	PublishDates string
}

type NewsPage struct {
	Title string
	News  []News
}

func jsonDataHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll((resp.Body))
	resp.Body.Close()

	fmt.Println("Fetching Json")

	var s NewsIndex
	error := xml.Unmarshal(bytes, &s)

	if error != nil {
		fmt.Println(error)
		return
	}

	jsonData, error := json.Marshal(s.NewsItems)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

func tableDataHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll((resp.Body))
	resp.Body.Close()

	fmt.Println("Fetching data")

	var s NewsIndex
	error := xml.Unmarshal(bytes, &s)

	if error != nil {
		fmt.Println(error)
		return
	}

	page := NewsPage{Title: "Cool Bussiness News", News: s.NewsItems}
	fmt.Println(page)
	htmlTemplate, error := template.ParseFiles("src/github.com/Learning-GoLang/TableDataTemplate.html")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(htmlTemplate.Execute(w, page))
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	fmt.Println("Fetching all data")
	http.HandleFunc("/jsondata", jsonDataHandler)
	http.HandleFunc("/tabledata", tableDataHandler)
	fmt.Println("FInised Fetching data")
	http.ListenAndServe(":8007", nil)
	fmt.Println("Serving data")
}
