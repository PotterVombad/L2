package pkg

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

type telnet struct {
	address string
	timeout time.Duration
}

func NewTelnet() (*telnet, error) {
	// парсим аргументы командной строки
	host := flag.String("host", "", "IP or Host")
	port := flag.Int("port", 0, "Port")
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()

	// проверяем, что хост и порт были переданы
	if *host == "" || *port == 0 {
		return nil, fmt.Errorf("give host and port to connect")
	}

	// формируем адрес для подключения
	address := fmt.Sprintf("%s:%d", *host, *port)

	return &telnet{
		address: address,
		timeout: *timeout,
	}, nil
}

func (t *telnet) Run() error {

	// подключаемся к серверу
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("smth with connection to server, %v", err)
	}
	defer conn.Close()

	errChan := make(chan error)
	sign := make(chan os.Signal, 1)
	// биндим обработчик сигнала
	signal.Notify(sign, os.Interrupt)

	// запускаем горутину для чтения данных из сокета и вывода в STDOUT
	go func(errChan chan error) {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				errChan <- fmt.Errorf("connection was closed, %v", err)
			}
			fmt.Print(string(buf[:n]))
		}
	}(errChan)

	// запускаем горутину для чтения данных из STDIN и отправляем их в сокет
	go func(errChan chan error) {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				errChan <- fmt.Errorf("smth with reading from socket, %v", err)
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				errChan <- fmt.Errorf("smth with writing to socket, %v", err)
			}
		}
	}(errChan)

	// ждем выхода времени, ошибки или закрытия ручками
	select {
	case c := <-sign:
		log.Println("Catch signal:", c)
		return nil
	case err := <-errChan:
		return err
	}
}
