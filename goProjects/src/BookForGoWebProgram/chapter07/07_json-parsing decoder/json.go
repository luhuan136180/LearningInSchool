package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("src/BookForGoWebProgram/chapter07/07_json-parsing decoder/post.json")
	if err != nil {
		fmt.Println("Error opeing JSON file:", err)
		return
	}
	defer jsonFile.Close()

	//
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR  decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}

}
