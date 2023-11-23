package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Определение флагов
	number := flag.String("f", "", "coloumn number")
	interp := flag.String("d", "\t", "another interpreter")
	delim := flag.Bool("s", false, "only with delim")
	flag.Parse()

	// Разбор входных строк
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// проверяем, что если есть флаг на разделитель, то пропускаем строки без него
		if *delim && !strings.Contains(line, *interp) {
			continue // Пропускаем строки без разделителя
		}

		fields := strings.Split(line, *interp)

		// вывод выбранных колонок
		if *number != "" {
			coloumns := parseColoumns(*number, len(fields))
			for _, field := range coloumns {
				fmt.Print(fields[field-1])
				if field < len(fields) {
					fmt.Print(*interp)
				}
			}
			fmt.Println()
		} else {
			// если не указаны колонки, просто выводим строку как есть
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения входных данных:", err)
	}
}

func parseColoumns(fieldsStr string, length int) []int {
	var fields []int
	fieldsStr = strings.TrimSpace(fieldsStr)
	fieldsArr := strings.Split(fieldsStr, ",")
	for _, f := range fieldsArr {
		num, err := fmt.Sscanf(f, "%d")
		if err != nil {
			fmt.Printf("%s is not an integer", f)
			continue
		}
		if num < 1 || num > length -1 {
			fmt.Printf("table does not include coloumn: %v", num)
			continue
		}
		fields = append(fields, num)
	}
	sort.Ints(fields)
	return fields
}
