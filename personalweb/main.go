package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index", nil)
}
