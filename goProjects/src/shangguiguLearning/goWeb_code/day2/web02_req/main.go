package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shangguiguLearning/goWeb_code/day1/web01_db/model"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//获取请求参数

	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.RawQuery)
	fmt.Fprintln(w, "你发送的请求头的所有信息：", r.Header) //获取请求头中的所有信息

	//type Header map[string][]string
	//获取请求头中的某个具体属性的值，如获取 Accept-Encoding 的值:方法一：
	fmt.Fprintln(w, "请求头中Accpet——Encoding的属性值是：", r.Header["Accept-Encoding"])
	//方法二：func (h Header) Get(key string) string
	fmt.Fprintln(w, "请求头中Accpet——Encoding的属性值是：", r.Header.Get("Accept-Encoding"))

	// //获取请求体重内容的长度
	// len := r.ContentLength
	// //创建byte切片
	// body := make([]byte, len)
	// //将请求体中的内容读到body中
	// r.Body.Read(body)
	// //在浏览器中显示请求体中的内容
	// fmt.Fprintln(w, "请求体中的内容有：", string(body))

	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面

	//	只能在  enctype="application/x-www-form-urlencoded"  编码下执行
	//r.ParseForm()
	//fmt.Fprintln(w, "请求参数有", r.Form)
	//fmt.Fprintln(w, "..POST请求的form表单中的请求参数有：", r.PostForm)

	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	//一下两种获取值的方法可以在  enctype="multipart/form-data"   编码 下执行
	//也可以在  enctype="application/x-www-form-urlencoded" 编码下执行
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))

	//5.4.4MultipartForm 字段
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
}
func testJsonRes(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("shoudao请求"))

	////给客户端响应一个 HTML 页面
	//html := `<html>
	//			<head>
	//				<title>测试响应内容为网页</title>
	//				<meta charset="utf-8"/>
	//			</head>
	//				<body>
	//					我是以网页的形式响应过来的！
	//				</body>
	//			</html>`
	//w.Write([]byte(html))

	//给客户端响应 JSON 格式的数据
	//设置响应内容的类型
	w.Header().Set("content-Type", "application/json")
	//创建User
	user := model.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}

func testRedire(w http.ResponseWriter, r *http.Request) {
	//设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/hello", handler)

	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedire)
	http.ListenAndServe(":8080", nil)
}
