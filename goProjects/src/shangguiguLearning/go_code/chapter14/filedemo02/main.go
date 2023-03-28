package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//读取文件的内容并显示在终端(带缓冲区的方式)
	//适用于文件较大的情况
	file,err:=os.Open("d:/test.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	//当函数退出时，要及时的关闭file
	defer file.Close()//要及时关闭file句柄，否则会有内存泄漏.

	// 创建一个 *Reader  ，是带缓冲的
	/*
		const (
		defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader:=bufio.NewReader(file)
	for{
		str,err:=reader.ReadString('\n')
		if err==io.EOF{
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束...")
}
