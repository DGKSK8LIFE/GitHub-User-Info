package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var html *template.Template

func init() {
	html = template.Must(template.ParseGlob("main.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		uInfo, err := http.Get("https://api.github.com/users/" + username)
		stringUinfo, err := ioutil.ReadAll(uInfo.Body)
		if err != nil {
			panic(err)
		}
		html.ExecuteTemplate(w, "main.html", nil)
		if len(username) != 0 {
			fmt.Fprintln(w, string(stringUinfo))
		}

	})
	http.ListenAndServe(":8080", nil)
}
