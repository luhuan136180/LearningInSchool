package main
import (
	"fmt"
)
type Monkey struct{
	Name string
}

func (this *Monkey) Climbing() {
	fmt.Println("huipashu")
}
type BirdAble interface {
	Flying()
}
type LittleMonkey struct{
	Monkey
}
func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name,"huifei")
}
func main() {
	money :=LittleMonkey{
		Monkey{
		Name: "wukong",
	},
	}
	money.Flying()
	var a BirdAble = &money
	a.Flying()
}

