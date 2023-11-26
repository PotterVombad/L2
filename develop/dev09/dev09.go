package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	urlFlag := flag.String("url", "", "URL of the website to download")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Please give URL using the -url flag")
		return
	}

	err := download(*urlFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func download(urlStr string) error {

	//получаем ответ
	response, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	fileName := "site_" + parsedURL.Host

	//проверяем статус код
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("connection failed. Status code: %d", response.StatusCode)
	}

	// создаем файл для записи полученных байт
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("does not create file: %v", err)
	}
	defer file.Close()

	// копируем тело сайта в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("does not copy site to file: %v", err)
	}
	return nil
}
