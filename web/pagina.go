package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/panel",panel)
	http.HandleFunc("/botones",botones)
	http.HandleFunc("/perfil",perfil)
	http.HandleFunc("/alarmas",alarmas)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mainmenu.gohtml", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func panel(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "panel.gohtml", nil)
}

func perfil(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "perfil.gohtml", nil)
}

func botones(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "botones.gohtml", nil)
}

func alarmas(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "alarmas.gohtml", nil)
}
