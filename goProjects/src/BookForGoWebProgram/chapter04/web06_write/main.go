package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello ,World,this is web06_write</h1></body>
</html>`
	w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service,try next door!by nhx")
}
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302) //执行该方法后禁止对首部进行修改，
	//所以，set 方法要在 writeheader 之前

}

//返回json格式
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()

}
