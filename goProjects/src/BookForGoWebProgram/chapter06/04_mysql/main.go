package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_web program_01")
	if err != nil {
		panic(err)
	}

}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id , content,author from posts limit ?", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id ,content,author from posts where id= ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) { //对mysql修改版
	statement := "insert into posts (content,author) values (?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.Content, post.Author)

	statement2 := "select max(id) from posts"
	stmt2, err := Db.Prepare(statement2)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt2.QueryRow().Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?,author = ? where id=?", post.Content, post.Author, post.Id)

	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}

func main() {
	post := Post{Content: "hello world!", Author: "mhx"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)
	////
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	//
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()
	//
	posts, _ := Posts(10)
	fmt.Println(posts)
	//
	readPost.Delete()
}
