package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

)


//func Dial(network, address string) (Conn, error)
func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Println("rclient dial err=",err)
		return
	}
	fmt.Println("CONN 成功",conn )
	//功能一：客户端可以发送单行数据，然后就退出
	reader:=bufio.NewReader(os.Stdin)

	for{
		line,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println("ReadString err=",err)
		}
		//如果用户输入的是 exit就退出
		line2 := strings.Trim(line, " \r\n")
		if line2 =="exit"{
			fmt.Println("客户端退出..")
			break
		}
		// Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
		// Write(b []byte) (n int, err error)
		n,err:=conn.Write([]byte(line))
		if err != nil{
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("客服端发送了%d字节的数据 \n",n)
	}
}