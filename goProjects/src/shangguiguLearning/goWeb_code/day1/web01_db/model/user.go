package model

import (
	"fmt"
	"shangguiguLearning/goWeb_code/day1/web01_db/utils"
)

//User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	//预编译(可以省略)
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编辑出现异常", err)
		return err
	}
	//执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@163.con")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}

	return nil
}

//AddUser2 添加User的方法二
func (user *User) AddUser2() error {
	//1.写sql语句
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	//2.执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "666666", "admin2@sina.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//查询：根据用户的 id 从数据库中获取一条记录
func (user *User) GetUserById(id int) (RTuser *User, err error) {
	sqlStr := "select id , username, password,email from user where id=?"
	//执行sql
	//func (db *DB) QueryRow(query string, args ...interface{}) *Row
	//QueryRow执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
	row := utils.Db.QueryRow(sqlStr, id)
	//type Row struct {
	//	// One of these two will be non-nil:
	//	err  error // deferred error for easy chaining
	//	rows *Rows
	//}

	//声明三个变量
	var username string
	var password string
	var email string
	//func (rs *Rows) Scan(dest ...interface{}) error
	//
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	RTuser = &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return RTuser, err
}

//GetUsers 获取数据库中所有的数据
func (user *User) GetUsers() (users []*User, err error) {
	sqlStr := "select id ,username,password,email from user"
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	//Query执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err

	}
	//Next准备用于Scan方法的下一行结果。如果成功会返回真，
	//如果没有下一行或者出现错误会返回假。Err应该被调用以区分这两种情况。
	//每一次调用Scan方法，甚至包括第一次调用该方法，都必须在前面先调用Next方法。
	for rows.Next() { //func (rs *Rows) Next() bool
		var userId int
		var username string
		var password string
		var email string
		err = rows.Scan(&userId, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		user := &User{
			ID:       userId,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, user)

	}
	return users, nil
}
