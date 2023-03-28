package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web01/tmpl.html")
	//Must 函数会自动报错
	t := template.Must(template.ParseFiles("src/BookForGoWebProgram/chapter05/web01/tmpl.html"))
	//当编译成main.exe 后 以及在终端中go run 时可以使用相对路径，直接用main.go 运行时需要绝对路径(或来自根目录路径)？
	t.Execute(w, "Hello World--by mhx web01 test42")
}
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/BookForGoWebProgram/chapter05/web01/tmpl.html",
		"src/BookForGoWebProgram/chapter05/web01/temp2.html")
	t.Execute(w, "Hello World!")
}
func main() {
	//wd, _ := os.Getwd()
	//fmt.Println(wd)
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)

	server.ListenAndServe()
}
