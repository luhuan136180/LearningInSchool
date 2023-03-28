package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 1; j < 10; j++ {
			sum++
		}
	}
	fmt.Println(sum)
}
