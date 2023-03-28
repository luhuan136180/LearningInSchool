package main

import "fmt"

func main() {
	var map1 = map[string]int{}
	map1["psw1"] = 123
	map1["psw2"] = 456
	map1["psw3"] = 789
	delete(map1, "pws2")
	map1["pws4"] = 135
	fmt.Println(map1)
}
