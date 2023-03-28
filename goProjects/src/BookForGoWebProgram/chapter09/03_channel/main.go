package main

import (
	"fmt"
	"time"
)

func PrintNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 1)
		fmt.Printf("%d", i)
	}
	w <- true
}
func PrintLetters2(w chan bool) {
	for i := 'A'; i <= 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	w <- true
}
func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go PrintNumbers2(w1)
	go PrintLetters2(w2)
	<-w1
	<-w2
}
