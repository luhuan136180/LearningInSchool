package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web08_content aware/tmpl.html")
	content := `I asked :<i>"What's up?"</i>`
	t.Execute(w, content)
}
func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe("localhost:8080", nil)
}
