package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web03_range array/tmpl.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web03_range array/tmpl2.html")
	//daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, nil)
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)

	server.ListenAndServe()
}
