package main
import (
	"fmt"
	"time"
)

func main() {
	now :=time.Now()
	fmt.Printf("now = %v,now type=%T",now,now)

	fmt.Printf("年=%v \n",now.Year())
	fmt.Printf("yue = %v \n",now.Month())
	fmt.Printf("yue = %v \n",int(now.Month()))
	fmt.Printf("ri = %v \n",now.Day())
	fmt.Printf("hour = %v \n",now.Hour())
	fmt.Printf("minute = %v \n",now.Minute())
	fmt.Printf("second = %v \n",now.Second())

	//格式化日期时间

	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v\n", dateStr)
	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	fmt.Printf(now.Format("2006"))
	fmt.Println()

	//需求，每隔1秒中打印一个数字，打印到100时就退出
	//需求2: 每隔0.1秒中打印一个数字，打印到100时就退出
	i:=0
	for{
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond*100)
		if i == 100{
			break
		}
	}

	//Unix和UnixNano的使用:时间戳用来获取随机数字
	fmt.Printf("unix时间戳=%v ,unixNano时间戳=%v",now.Unix(),now.UnixNano())

}