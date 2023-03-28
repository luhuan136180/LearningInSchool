package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	m := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, m)
	n := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, n)

}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
