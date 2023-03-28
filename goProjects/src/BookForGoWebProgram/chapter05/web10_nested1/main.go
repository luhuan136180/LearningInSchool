package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web10_nested1/layout.html")
	//t.Execute(w, "1")  该方法无法在此情况下运行
	t.ExecuteTemplate(w, "layout", "2")
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
