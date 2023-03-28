package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form["hello"]) //返回URL和表单中的键值对
	fmt.Fprintln(w, r.PostForm)      //只返回表单中的键值对
	//主义：PostForm 函数只支持enctype="application/x-www-form-urlencoded" 编码

	// 获取

}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
