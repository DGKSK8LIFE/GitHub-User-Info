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
	PublicRepos int
	PublicGists int
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
		getJSON("https://api.github.com/users/"+username, user)
		html.ExecuteTemplate(w, "main.html", nil)
		if len(username) != 0 {
			fmt.Fprintf(w, "<h3 style=\"text-align: center; \">Name: %s<br>Login: %s<br>ID: %d<br>Company: %s<br>Blog: %s<br>Location: %s<br>Email: %s<br>Hireable: %t<br>Bio: %s<br>Public Repositories: %d<br>Public Gists: %d<br>Followers: %d<br>Following: %d<br></h3>", user.Name, user.Login, user.ID, user.Company, user.Blog, user.Location, user.Email, user.Hireable, user.Bio, user.PublicRepos, user.PublicGists, user.Followers, user.Following)
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
