package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	err = response.Validate()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	time := time.Now().Add(response.ClockOffset).Format("15:04:05")
	fmt.Println(time)
}
