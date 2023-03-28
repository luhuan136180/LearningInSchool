package main

import "fmt"

//⼀球从100m⾼度⾃由落下，每次落地后反跳回原来⾼度的⼀半，
//再落下，求它在第10次落地时，共经过多少⽶？第10次反弹多⾼？
func main() {
	h := 10.0
	s := 0.0
	for i := 1; i <= 10; i++ {
		s += h
		h = h / 2.0
		s += h
	}
	s = s - h
	fmt.Printf("共经过%v米，第十次反弹时高度为：%v", s, h)
}
