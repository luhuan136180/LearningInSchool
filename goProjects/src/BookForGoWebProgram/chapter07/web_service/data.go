package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	Db  *sql.DB
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_web program_02")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	sqlStr := "select id,content,author from posts where id =?"
	err = Db.QueryRow(sqlStr, id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
func (post *Post) create() (err error) {
	fmt.Println(post)
	sqlStr := "insert into posts (content,author) values (?,?)"
	_, err = Db.Exec(sqlStr, post.Content, post.Author)
	sqlStr2 := "select max(id) from posts"
	err = Db.QueryRow(sqlStr2).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	sqlStr := "update posts set content=?,author=? where id=?"
	_, err = Db.Exec(sqlStr, post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id=?", post.Id)
	return
}
