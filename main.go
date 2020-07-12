package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prathyushnallamothu/project1/structs"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/", homehandler)
	http.ListenAndServe(":8080", m)
}
func homehandler(w http.ResponseWriter, r *http.Request) {
	var person1 structs.Person
	person1.Name = "prathyush"
	person1.Email = "prathyushbablu@gmail.com"
	tpl.ExecuteTemplate(w, "home.html", person1)
}
