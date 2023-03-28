package main

import (
	"fmt"
	"net/http"
)
type MyHandler struct{

}

func(h	*MyHandler)	ServeHTTP(w	http.ResponseWriter,r *http.Request) {

fmt.Fprintln(w, "正在通过处理器处理你的请求~")

}


func main() {

	//myHandler := MyHandler{}
	////调用处理器
	//http.Handle("/", &myHandler)
	//http.ListenAndServe(":8080", nil)
	mux:=http.NewServeMux()
	mux.Handle("/",&MyHandler{})
	http.ListenAndServe(":8080",mux)
}