package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/read-cookies", readCookies)
	http.HandleFunc("/write-cookies", writeCookies)
	http.HandleFunc("/dump-cookies", dumpCookies)
	http.HandleFunc("/count", visitCounter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func visitCounter(writer http.ResponseWriter, req *http.Request) {
	count, error := req.Cookie("count")

	if error != nil {
		incrementVisitCount(writer, "1")
		fmt.Fprintln(writer, "Welcome, this is your first visit")
	} else {
		integer, error := strconv.Atoi(count.Value)
		if error != nil {
			fmt.Fprintln(writer, "A weird error occured")
		} else {
			incrementVisitCount(writer, strconv.Itoa(integer+1))
			fmt.Fprintln(writer, "Welcome back, this is visit no: ", count)
		}
	}
}

func incrementVisitCount(writer http.ResponseWriter, count string) {
	http.SetCookie(writer, &http.Cookie{
		Name:  "count",
		Value: count,
	})
}

func readCookies(writer http.ResponseWriter, req *http.Request) {
	cookie1, error := req.Cookie("my-custom-cookie")

	if error != nil {
		log.Println(error)
	} else {
		fmt.Fprintln(writer, "Found a cookie:", cookie1)
	}

	invalidCookie, error1 := req.Cookie("my-custom-cookie23")

	if error1 != nil {
		log.Println(error1)
	} else {
		fmt.Fprintln(writer, "Found a cookie:", invalidCookie)
	}
}

func writeCookies(writer http.ResponseWriter, req *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:  "my-custom-cookie",
		Value: "Hamadan or Hamedan (Persian: همدان, Hamedān) is the capital city of Hamadan Province of Iran. At the 2019 census, its population was 783,300 in 230,775 ...\n",
	})

	fmt.Println(writer, "Cookie Written - Check your browser")
	fmt.Println(writer, "In chrome go to: dev tolls / application / cookies")
}

func dumpCookies(writer http.ResponseWriter, req *http.Request) {
	cookieToDelete, error := req.Cookie("count")
	if error != nil {
		fmt.Println(writer, "An error occurred")
		fmt.Println(writer, error.Error(), http.StatusInternalServerError)
	} else {
		cookieToDelete.MaxAge = -1
		http.SetCookie(writer, cookieToDelete)
		fmt.Println(writer, "Counter cookie deleted")
		http.Redirect(writer, req, "/count", http.StatusSeeOther)
	}
}
