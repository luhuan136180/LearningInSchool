package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"shangguiguLearning/go_code/chatroom/common/message"
)
//这里将方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte//传输时用以缓冲,里面的内容在随时进行覆写操作，

}
func (this *Transfer)WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}



func (this *Transfer)ReadPkg()(mes message.Message, err error) {
	//我们将读取数据包封装成一个函数
	//buf:=make([]byte,8096)
	fmt.Println("读取客户端发送的数据。。。")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞
	_,err=this.Conn.Read(this.Buf[:4])
	if err!=nil{
		fmt.Println("conn.Read(bytes) err", err)
		return
	}

	fmt.Println("读到的buf=",this.Buf[:4])
	//将【】byte型：buf 转换为一个uint32数据类型
	var pkgLen uint32
	pkgLen=binary.BigEndian.Uint32(this.Buf[:4])//获取长度

	//根据pkglen 读取消息内容
	n,err:=this.Conn.Read(this.Buf[:pkgLen])//接收消息
	//fmt.Println("buf=",buf)//发过来的时数字串
	fmt.Printf("n=%d,pkglen=%d \n",n,pkgLen)
	if uint32(n) != pkgLen||err!=nil{//比对读取的信息长度是否与pkglen相同

		fmt.Println("readpkg body error")
		return
	}
	//对信息的比对完成，现在将信息反序列化，获得信息体message
	//位置原因，反序化是接收需要用地址（指针方式接收）
	err=json.Unmarshal(this.Buf[:pkgLen],&mes)//将反序化后的信息写入 mes中
	if err!=nil{
		fmt.Println("json.Unmarsha err=",err)
		return
	}
	return
	//
}