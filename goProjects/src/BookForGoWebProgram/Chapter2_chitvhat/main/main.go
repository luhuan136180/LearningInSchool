package main

import (
	"BookForGoWebProgram/Chapter2_chitvhat/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")

		}
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)
	//
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signup_account)
	//mux.HandleFunc("/authenticate", authentiacate)
	//
	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("thread/read", readThread)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
