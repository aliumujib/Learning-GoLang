package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var AdminRole = "ADMIN"
var UserRole = "USER"

type User struct {
	FName          string
	LName          string
	HashedPassword string
	Role           string
}

var sessionDb map[string]string
var userDb map[string]User

func init() {
	sessionDb = map[string]string{}
	userDb = map[string]User{}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/user-list", UserList)
	http.HandleFunc("/drinks-menu", DrinksMenu)
	http.HandleFunc("/logout", Logout)

	http.ListenAndServe(":8080", nil)
}

func DrinksMenu(writer http.ResponseWriter, request *http.Request) {

}

func Register(writer http.ResponseWriter, request *http.Request) {
	//make templates
	//get user struct
	//generate uuid
	//store user in db
	//redirect to drinks-menu
	hashedPassword := request.FormValue("password")
	var role string

	if request.FormValue("admin") == "on" {
		role = AdminRole
	} else {
		role = UserRole
	}

	user := User{
		request.FormValue("first"),
		request.FormValue("last"),
		hashedPassword,
		role,
	}

	fmt.Println(user)

	tpl := template.Must(template.ParseGlob("sessions/templates/*"))

	error := tpl.ExecuteTemplate(writer, "register-page.html", nil)
	if error != nil {
		http.Error(writer, error.Error(), 500)
	}
}

func UserList(writer http.ResponseWriter, request *http.Request) {

}

func Logout(writer http.ResponseWriter, request *http.Request) {

}

func Login(writer http.ResponseWriter, request *http.Request) {

}

func Index(writer http.ResponseWriter, request *http.Request) {

}
