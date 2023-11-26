package pkg

import (
	"flag"
	"fmt"
)

type Flags struct {
	coloumn       int
	numeric       bool
	reverse       bool
	unique        bool
	month         bool
	whitespace    bool
	checkSort     bool
	numericSuffix bool
	path          string
}

func NewFlags() (*Flags, error) {
	key := flag.Int("k", 1, "указание колонки для сортировки")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	month := flag.Bool("M", false, "сортировать по названию месяца")
	whitespace := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSort := flag.Bool("c", false, "проверять отсортированы ли данные")
	numericSuffix := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")


	flag.Parse()

	if *key < 1 {
		return nil, fmt.Errorf("invalid coloumn for sort")
	}

	filePath := flag.Arg(0) // Путь к файлу с данными
	if *key < 1 || filePath == "" {
		return nil, fmt.Errorf("invalid file name ")
	}

	if *checkSort {
		if *numeric || *reverse || *unique || *month || *whitespace || *numericSuffix {
			return nil, fmt.Errorf("check sort must be alone")
		}
	}

	input := Flags{
		coloumn:       *key,
		numeric:       *numeric,
		reverse:       *reverse,
		unique:        *unique,
		month:         *month,
		whitespace:    *whitespace,
		checkSort:     *checkSort,
		numericSuffix: *numericSuffix,
		path:          filePath,
	}

	return &input, nil
}
