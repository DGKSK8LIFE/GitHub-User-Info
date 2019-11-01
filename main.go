package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name        string
	Login       string
	ID          int
	Company     string
	Blog        string
	Location    string
	Email       string
	Hireable    bool
	Bio         string
	PublicRepos string
	PublicGists string
	Followers   int
	Following   int
}

var html *template.Template

func init() {
	html = template.Must(template.ParseGlob("main.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		user := &UserInfo{}
		userJSON := getJSON("https://api.github.com/users/"+username, user)
		html.ExecuteTemplate(w, "main.html", nil)
		if len(username) != 0 {
			fmt.Fprintf(w, "Info: %v", userJSON)
		}

	})
	http.ListenAndServe(":8080", nil)
}

var site = &http.Client{}

func getJSON(url string, target interface{}) error {
	r, err := site.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
