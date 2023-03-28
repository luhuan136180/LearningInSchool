package main

import (
	"fmt"
	"os"
	_ "os/exec"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

}
