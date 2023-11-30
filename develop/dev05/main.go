package main

import (
	"dev05/pkg"
	"fmt"
)

func main() {
	grep, err := pkg.NewGrep()
	if err != nil {
		fmt.Println(err)
	}
	err = grep.Run()
	if err != nil {
		fmt.Println(err)
	}
}