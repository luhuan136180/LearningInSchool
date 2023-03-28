package main

import "fmt"

func main() {
	//Var    map变量名   map[keytype]valuetype
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	//同一个key赋值会覆盖，但value可以重复
	a = make(map[string]string,4)
	a["no1"]="ssss"
	a["no2"]="aaa"
	fmt.Println(a)
	a["no1"]="2sss"
	fmt.Println(a)
	a["n03"]="www"
	fmt.Println(a)
	a["929a"]="xxx"
	fmt.Println(a)
}
