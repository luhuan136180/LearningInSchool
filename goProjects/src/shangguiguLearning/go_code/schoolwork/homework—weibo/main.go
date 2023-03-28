package main

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"

)

func main() {

	db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/classworkforgo_weibo")
	defer db.Close()	//关闭数据库
	err:=db.Ping()
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}


}