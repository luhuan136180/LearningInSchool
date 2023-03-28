package main
import (
	"bufio"
	"fmt"
	"os"
)
//打开一个存在的文件中，将原来的内容
//os.O_TRUNC覆盖
//成新的内容 10 句 "你好，尚硅谷!"
func main() {
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath:="d:/abc.txt"
	file,err :=os.OpenFile(filePath,os.O_WRONLY|os.O_TRUNC,0666)
	if err!=nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//及时关闭file语句
	defer file.Close()
	//准备写入5句
	str:="你好，尚硅谷！~\r\n"
	//写入时，使用带缓存的 *Writer
	writer:=bufio.NewWriter(file)//读入相应file文件
	for i:=0;i<10;i++{
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}