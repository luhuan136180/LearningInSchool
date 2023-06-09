package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("src/BookForGoWebProgram/chapter05/web11_nested2/layout.html",
			"src/BookForGoWebProgram/chapter05/web11_nested2/red_hello.html")

	} else {
		t, _ = template.ParseFiles("src/BookForGoWebProgram/chapter05/web11_nested2/layout.html",
			"src/BookForGoWebProgram/chapter05/web11_nested2/blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func process2(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("src/BookForGoWebProgram/chapter05/web11_nested2/layout.html",
			"src/BookForGoWebProgram/chapter05/web11_nested2/red_hello.html")

	} else {
		t, _ = template.ParseFiles("src/BookForGoWebProgram/chapter05/web11_nested2/layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)

	server.ListenAndServe()
}
