package main

import (
	"dev03/pkg"
	"fmt"
)

func main() {
	sorter, err := pkg.NewSorter()
	if err != nil {
		fmt.Println(err)
		return
	}
	if sorter == nil {
		return
	}
	err = sorter.Run()
	if err != nil {
		fmt.Println(err)
		return
	} 
	fmt.Println("sorted successfully")
}
