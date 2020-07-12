package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}
func main(){
	m:=mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/",homehandler)
	http.ListenAndServe(":8080",m)
}
func homehandler(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"home.html",nil)
}