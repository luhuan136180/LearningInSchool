package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formatDate, //将名字fdate映射至 formatDate 函数
	}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("src/BookForGoWebProgram/chapter05/web07_custom function/tmpl.html")
	t.Execute(w, time.Now())
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
