package main

import "fmt"

func main() {
	//Var    map变量名   map[keytype]valuetype
	var a map[string]string
	a = make(map[string]string,4)
	a["no1"]="ssss"
	a["no2"]="aaa"
	fmt.Println(a)
	a["no1"]="2sss"
	fmt.Println(a)
	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	var b map[string]string=make(map[string]string,10)
	b["ww"]="qq"
	b["ss"]="aa"
	fmt.Println(b)

	//第三种方式
	heroes := map[string]string{
		"hero1" : "宋江",
		"hero2" : "卢俊义",
		"hero3" : "吴用",
	}
	heroes["hero4"] = "林冲"
	fmt.Println("heroes=", heroes)

	//案例
	/*
		课堂练习：演示一个key-value 的value是map的案例
		比如：我们要存放3个学生信息, 每个学生有 name和sex 信息
		思路:   map[string]map[string]string

	*/
	studentMap :=make(map[string]map[string]string)

	studentMap["student01"]=make(map[string]string,3)
	studentMap["student01"]["name"]="tom"
	studentMap["student01"]["sex"]="男"
	studentMap["student01"]["adress"]="s"

	studentMap["stu02"] =  make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"

	fmt.Println(studentMap)
	fmt.Println(studentMap["student01"])
	fmt.Println(studentMap["stu02"]["name"])

}