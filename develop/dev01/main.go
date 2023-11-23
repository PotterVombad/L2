package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	//запрашиваем время с помощью библиотеки
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	//отрабатываем ошибку при наличии
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	//печатаем время
	_, err = fmt.Println(time.Format("15:04:0523"))
	//отрабатываем ошибку при наличии
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
