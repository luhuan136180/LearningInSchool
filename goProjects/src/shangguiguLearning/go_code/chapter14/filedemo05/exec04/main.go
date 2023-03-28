package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//打开一个存在的文件，
//将原来的内容读出显示在终端，
//并且追加 5 句"hello,北京!"

func main() {
	//打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath := "d:/abc.txt"
	file,err:=os.OpenFile(filePath,os.O_RDWR|os.O_APPEND,0666)
	if err!=nil{
		fmt.Printf("Open  file err=%v \n",err)
		return
	}
	defer file.Close()
	reder:=bufio.NewReader(file)
	for{
		str,err:=reder.ReadString('\n')
		if err==io.EOF{////如果读取到文件的末尾
			break
		}
		fmt.Print(str)
	}
	str :="HELLO,BEIJIGN!\r\n"
	writer:=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()

}

