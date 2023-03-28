package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first-cookie",
		Value:    "Go web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Set("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2) //注：使用指针

}
func getCookie(w http.ResponseWriter, r *http.Request) {
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)
	c1, err := r.Cookie("first_cookie") //获取指定的cookie
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}

	cs := r.Cookies() //返回包含了所有cookie的切片
	fmt.Fprintln(w, c1)
	fmt.Println("得到的cookie为:", c1)
	fmt.Fprintln(w, cs)
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
