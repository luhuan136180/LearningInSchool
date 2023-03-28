package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form["hello"])//返回URL和表单中的键值对
	//fmt.Fprintln(w,r.PostForm)//只返回表单中的键值对
	////主义：PostForm 函数只支持enctype="application/x-www-form-urlencoded" 编码

	// 获取 "multipart/form-data" 编码格式中的 表单数据
	// 我们需要使用  ParseMultipartForm 方法 和
	//  MultipartForm 字段
	//"multipart/form-data" 编码格式 会主动将数据存入  MultipartForm 字段
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
	//PostFormValue  方法只能输出一对值
	fmt.Fprintln(w, r.PostFormValue("hello"))
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
