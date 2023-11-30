package main

import (
	"dev10/pkg"
	"fmt"
)

func main() {
	telnet, err := pkg.NewTelnet()
	if err != nil {
		fmt.Printf("mistace with parsing adress: %v", err)
	}
	err = telnet.Run()
	if err != nil {
		fmt.Println(err)
	}
}