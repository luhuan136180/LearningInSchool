package main

import (
	"fmt"
	"sync"
	"time"
)

func printNUmbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 1)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNUmbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
