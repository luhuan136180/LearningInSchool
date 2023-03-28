package main
import (
	"fmt"
	"io/ioutil"

)
//编程一个程序，将一个文件的内容，写入到另外一个文件。注：这两个文件已经存在了.
//说明：使用 ioutil.ReadFile / ioutil.WriteFile 完成写文件的任务.
//代码实现：
func main() {
	//将d:/abc.txt 文件内容导入到  e:/kkk.txt
	//1. 首先将  d:/abc.txt 内容读取到内存
	//2. 将读取到的内容写入 e:/kkk.txt
	file1Path:="d:/abc.txt"
	file2Path:="d:/def.txt"
	data,err:=ioutil.ReadFile(file1Path)
	if err !=nil{
		//说明读取文件有错误
		fmt.Printf("READ File err=%v \n",err)
		return
	}
	err=ioutil.WriteFile(file2Path,data,0666)
	if err!=nil{
		fmt.Printf("write file error =%v \n",err)
	}
}