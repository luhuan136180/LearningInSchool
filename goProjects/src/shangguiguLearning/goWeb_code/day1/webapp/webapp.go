package main

import (
	"fmt"
	"net/http"
)

//创建一个处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"hello world!",r.URL.Path)
}
func main() {
	//HanderFunc(handler) 会将 handler 转换成一个实现了 Handler 接口的 HandlerFunc 类型
	http.HandleFunc("/",handler)

	http.ListenAndServe(":8080",nil)
}

