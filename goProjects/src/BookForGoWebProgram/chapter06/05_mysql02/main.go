package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}
type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var (
	Db *sql.DB
)

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_web program_02")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	sqlStr := "insert into comments (content,author,post_id) values (?,?,?)"
	_, err = Db.Exec(sqlStr, comment.Content, comment.Author, comment.Post.Id)
	sqlStr2 := "select max(id) from posts"
	err = Db.QueryRow(sqlStr2).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	sqlStr := "select id,content,author from posts where id=?"
	err = Db.QueryRow(sqlStr, id).Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("select id,content,author from comments where post_id=?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}
func (post *Post) Create() (err error) {
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

func main() {
	post := Post{
		Content: "hello world 05",
		Author:  "sau ---mhx",
	}
	post.Create()
	comment := Comment{
		Content: "Good post!",
		Author:  "joe smith",
		Post:    &post,
	}
	comment.Create()
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
