package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//ParseFiles 函数用于解析模板，对模板进行语法分析
	//t := template.Must(template.ParseFiles("src/BookForGoWebProgram/chapter05/web04_set dot/tmpl.html"))
	t := template.Must(template.ParseFiles("src/BookForGoWebProgram/chapter05/web04_set dot/tmpl2.html"))
	t.Execute(w, "hello")
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
