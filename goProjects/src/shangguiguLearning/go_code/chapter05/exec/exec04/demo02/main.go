package main
import "fmt"
func main(){
	var score int
	fmt.Println("shuru hcengji ")
	fmt.Scanln(&score)
	switch {
	case score >= 60 :
		fmt.Println("合格")
	case score < 60:
		fmt.Println("buhege")
	default:
		fmt.Println("shuruyouwu")
	}


}
