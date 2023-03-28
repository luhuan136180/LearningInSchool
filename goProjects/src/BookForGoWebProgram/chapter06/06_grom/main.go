package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    //`sql:"index"`
	CreatedAt time.Time
}

var (
	err error
	Db  *gorm.DB
)

func init() {
	//var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/go_web program_02?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	err = Db.AutoMigrate(&Post{}, &Comment{})

}
func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) // {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}

	////// Create a post
	Db.Create(&post)

	fmt.Println(post.Id)

	fmt.Println(post) // {1 Hello World! Sau Sheong [] 2015-04-13 11:38:50.91815604 +0800 SGT}

	//comment := Comment{Content: "good post 1", Author: "mhx"}
	//Db.Model(&post).Association("Comments").Append(comment)

	//var readPost Post
	//Db.Where("author=?","Sau Sheong").First(&readPost)
	//var comments []Comment
	//Db.Model(&readPost).
}
