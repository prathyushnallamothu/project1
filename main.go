package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prathyushnallamothu/project1/structs"
)

var tpl *template.Template
var x string
func init() {
	err:=godotenv.Load("./.env")
	if err!=nil{
		log.Fatal(err)
	}
	x=os.Getenv("MY_FAV_HERO")
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/", homehandler)
	fmt.Println(x)
	http.ListenAndServe(":8080", m)
}
func homehandler(w http.ResponseWriter, r *http.Request) {
	var person1 structs.Person
	person1.Name = "prathyush"
	person1.Email = "prathyushbablu@gmail.com"
	tpl.ExecuteTemplate(w, "home.html", person1)
}
