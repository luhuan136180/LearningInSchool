package main

import (
	"fmt"
	"time"
)

func printNumbersl() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
}
func printLettersl() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func print1() {
	printNumbersl()
	printLettersl()
}
func goPrint1() {
	go printNumbersl()
	go printLettersl()
}
func main() {
	goPrint1()
	time.Sleep(time.Second * 20)
}
