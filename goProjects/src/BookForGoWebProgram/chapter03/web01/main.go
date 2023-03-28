package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!everyone")
}
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &handler,
	}
	server.ListenAndServe() //该指令会进入无限for循环
	fmt.Println("hahah")
}
