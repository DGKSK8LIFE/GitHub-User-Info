package main

import (
	"html/template"
	"net/http"
)

var html *template.Template

func init() {
	html = template.Must(template.ParseGlob("main.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		html.ExecuteTemplate(w, "main.html", nil)
	})
	http.ListenAndServe(":8080", nil)
}
