package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func CopyFile(dstFileName string,srcFileName string)(written int64,err error){
	srcFile,err :=os.Open(srcFileName)
	if err!=nil{
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader:=bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile,err:=os.OpenFile(dstFileName,os.O_WRONLY|os.O_CREATE,0666)
	if err!=nil{
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//通过dstFile, 获取到 Writer
	writer:=bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer,reader)
}
func main(){
	srcFile:="d:/abc.jpg"
	dstFile:="d:/gotest/cdf.jpg"
	_,err:=CopyFile(dstFile,srcFile)
	if err==nil{
		fmt.Printf("拷贝完成\n")
	}else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}
