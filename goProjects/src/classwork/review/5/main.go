package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		n := 10
		if i > 10 {
			n = 100
		}
		if i > 100 {
			n = 1000
		}
		if i*i%n == i {
			fmt.Println(i)
		}
	}
}
