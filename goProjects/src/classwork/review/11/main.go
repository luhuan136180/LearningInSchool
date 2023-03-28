package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var sum int
	buf := make([]byte, 1024)
	f, _ := os.Open("words.txt")
	defer f.Close()
	n, _ := f.Read(buf)
	for _, sentence := range bytes.FieldsFunc(buf[:n], f1) {
		for _, words := range bytes.Fields(sentence) {
			if bytes.Contains([]byte(words), []byte("er")) {
				fmt.Println(string(words))
				sum++
			}
		}
	}
	fmt.Println("含有er的单词总数为", sum)
}

func f1(a rune) bool {
	return a == ',' || a == '.' || a == '!'
}