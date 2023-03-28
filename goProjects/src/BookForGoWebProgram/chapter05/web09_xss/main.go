package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web09_xss/tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}
func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web09_xss/form.html")
	t.Execute(w, nil)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}
