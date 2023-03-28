package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/BookForGoWebProgram/chapter05/web05_include/t1.html",
		"src/BookForGoWebProgram/chapter05/web05_include/t2.html"))
	t.Execute(w, "Hello World!")

}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
