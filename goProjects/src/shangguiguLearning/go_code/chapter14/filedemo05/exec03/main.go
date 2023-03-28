package main

import (
	"bufio"
	"fmt"
	"os"
)

//打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'
func main() {
	//打开一个存在的文件，在原来的内容
	//追加: os.0_APPEND
	//内容 'ABC! ENGLISH!'
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath:="d:/abc.txt"
	file,err :=os.OpenFile(filePath,os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	str := "ABC,ENGLISH!\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer:=bufio.NewWriter(file)
	for i:=0;i<19;i++{
		writer.WriteString(str)//将str的内容输入缓存区
	}
	writer.Flush()
}