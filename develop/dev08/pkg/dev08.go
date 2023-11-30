package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	scanner *bufio.Scanner
}

func NewShell() *Shell {
	scanner := bufio.NewScanner(os.Stdin)

	return &Shell{
		scanner: scanner,
	}
}

func (s *Shell) Run() {

	// в бесконечном цикле отрабатываем запросы
	for {
		fmt.Print("> ")
		s.scanner.Scan()
		input := s.scanner.Text()

		// проверка на команду выхода
		if input == "\\quit" {
			fmt.Println("quit shell")
			return
		}

		// разбиваем введенную строку на команды для конвейера
		commands := strings.Split(input, "|")

		var cmd *exec.Cmd
		var err error

		//проход по каждой команде в конвейере
		for _, command := range commands {
			command = strings.TrimSpace(command)
			args := strings.Fields(command)

			// Проверка на наличие команды
			if len(args) == 0 {
				continue
			}

			if args[0] == "fork" || args[0] == "exec" {
				cmd = exec.Command(args[1], args[2:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("fork/exec error: %v", err)
				}
				continue
			}

			switch args[0] {

			// обработка команды cd
			case "cd":
				if len(args) > 1 {
					err = os.Chdir(args[1])
					if err != nil {
						log.Printf("cd error: %v", err)
					}
				} else {
					log.Print("no argument for cd")
				}
				continue

			// обработка команды pwd
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					log.Printf("pwd error: %v", err)
				} else {
					fmt.Println(dir)
				}
				continue

			// обработка команды echo
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
				continue

			// обработка команды kill
			case "kill":
				if len(args) > 1 {
					pid := args[1]
					err = exec.Command("kill", pid).Run()
					if err != nil {
						log.Printf("kill error: %v", err)
					}
				} else {
					fmt.Println("unknown PID for kill.")
				}
				continue

			// обработка команды ps
			case "ps":
				cmd = exec.Command("ps")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("ps error: %v", err)
				}
				continue

			default:
				log.Printf("unknown command: %s", args[0])
			}
		}
	}
}
