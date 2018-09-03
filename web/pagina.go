package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/wawis",wawis)
	http.HandleFunc("/panel",panel)
	http.HandleFunc("/perfil",perfil)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mainmenu.html", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func wawis(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "wawis.html", nil)
}

func panel(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "panel.html", nil)
}

func perfil(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "perfil.html", nil)
}
