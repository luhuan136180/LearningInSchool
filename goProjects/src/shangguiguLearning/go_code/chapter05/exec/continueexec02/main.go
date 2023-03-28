package main
import "fmt"

func main()  {
	var money float64 = 100000
	var i int = 0
	for{
		if money >50000{
			money = 0.95*money
			i++
		}else if money<=50000{
			money -=1000
			i++
			if money<1000{
				break
			}
		}
	}
	fmt.Println("i=",i)
}
