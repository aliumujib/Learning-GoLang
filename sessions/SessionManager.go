package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var AdminRole = "ADMIN"
var UserRole = "USER"

var SessionIdKey = "SID"

type User struct {
	UserId         string
	FName          string
	LName          string
	HashedPassword []byte
	Role           string
}

var sessionDb map[string]string
var userDb map[string]User

var tpl *template.Template

func init() {
	sessionDb = map[string]string{}
	userDb = map[string]User{}
	tpl = template.Must(template.ParseGlob("sessions/templates/*"))
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
	if !isLoggedIn(request) {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	sessionId, error := request.Cookie(SessionIdKey)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	user := userDb[sessionId.Value]

	tplError := tpl.ExecuteTemplate(writer, "drinks-menu.html", user.FName)
	if tplError != nil {
		http.Error(writer, tplError.Error(), 500)
	}
}

func isLoggedIn(request *http.Request) bool {
	_, cookieError := request.Cookie(SessionIdKey)
	return cookieError == nil
}

func Register(writer http.ResponseWriter, request *http.Request) {
	if isLoggedIn(request) {
		http.Redirect(writer, request, "/drinks-menu", http.StatusSeeOther)
		return
	}

	if request.Method == http.MethodPost {

		unHashedPassword := request.FormValue("password")
		var role string

		if request.FormValue("admin") == "on" {
			role = AdminRole
		} else {
			role = UserRole
		}

		byteStream, error := bcrypt.GenerateFromPassword([]byte(unHashedPassword), bcrypt.MinCost)
		if error != nil {
			http.Error(writer, error.Error(), http.StatusInternalServerError)
			return
		}

		user := User{
			uuid.NewV4().String(),
			request.FormValue("first"),
			request.FormValue("last"),
			byteStream,
			role,
		}

		saveSessionData(writer, user)

		fmt.Println("Sessions: ", sessionDb)
		fmt.Println("Users: ", userDb)

		http.Redirect(writer, request, "/drinks-menu", http.StatusSeeOther)
		return
	}

	tplError := tpl.ExecuteTemplate(writer, "register-page.html", nil)
	if tplError != nil {
		http.Error(writer, tplError.Error(), 500)
	}
}

func saveSessionData(writer http.ResponseWriter, user User) {
	sessionId := uuid.NewV4()

	sessionDb[sessionId.String()] = user.FName
	userDb[user.FName] = user

	sidCookie := &http.Cookie{
		Name:  SessionIdKey,
		Value: sessionId.String(),
	}

	http.SetCookie(writer, sidCookie)
}

func UserList(writer http.ResponseWriter, request *http.Request) {
	if !isLoggedIn(request) {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	sessionId, error := request.Cookie(SessionIdKey)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	user, ok := userDb[sessionDb[sessionId.Value]]
	if !ok {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Current User: ", user)

	if user.Role != AdminRole {
		http.Error(writer, "You don't have access to this resource", http.StatusUnauthorized)
		return
	}

	userList := make([]User, 0, len(userDb))
	for _, val := range userDb {
		userList = append(userList, val)
	}

	tplError := tpl.ExecuteTemplate(writer, "user-list.html", userList)
	if tplError != nil {
		http.Error(writer, tplError.Error(), 500)
	}

}

func Logout(writer http.ResponseWriter, request *http.Request) {
	cookieToDelete, error := request.Cookie(SessionIdKey)
	if error != nil {
		fmt.Println(writer, "An error occurred")
		fmt.Println(writer, error.Error(), http.StatusInternalServerError)
	} else {
		cookieToDelete.MaxAge = -1
		http.SetCookie(writer, cookieToDelete)
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	//store user in db
	//save cookie
	//redirect to drinks-menu
	if isLoggedIn(request) {
		http.Redirect(writer, request, "/drinks-menu", http.StatusSeeOther)
	}

	if request.Method == http.MethodPost {

		userName := request.FormValue("first")
		unHashedPassword := request.FormValue("password")

		user, ok := userDb[userName]
		if !ok {
			http.Error(writer, "Username and/or password is incorrect", http.StatusUnauthorized)
			return
		}

		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(unHashedPassword))
		if err != nil {
			http.Error(writer, "Username and/or password is incorrect", http.StatusUnauthorized)
			return
		}

		saveSessionData(writer, user)

		http.Redirect(writer, request, "/drinks-menu", http.StatusSeeOther)
		return
	}

	tplError := tpl.ExecuteTemplate(writer, "login-page.html", nil)
	if tplError != nil {
		http.Error(writer, tplError.Error(), 500)
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {

}
