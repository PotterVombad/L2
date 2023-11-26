package pkg

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Sorter struct {
	flags *Flags
	data  []string
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// удаляет повторяющиеся строки
func makeUnique(data []string) []string {
	var result []string
	seen := make(map[string]struct{})
	for _, line := range data {
		_, ok := seen[line]
		if !ok {
			result = append(result, line)
			seen[line] = struct{}{}
		}
	}
	return result
}

func justSort(lines []string) []string {
	var low []string
	mapLowerLinesToLines := map[string]string{}

	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		low = append(low, strings.ToLower(line))
		mapLowerLinesToLines[lowerLine] = line
	}
	sort.Strings(low)

	var result []string
	for _, key := range low {
		result = append(result, mapLowerLinesToLines[key])
	}
	return result
}

func slicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func (s *Sorter) Run() ([]string, error) {

	if s.flags.month && s.flags.numeric || s.flags.month && s.flags.numericSuffix || s.flags.numeric && s.flags.numericSuffix {
		return nil, fmt.Errorf("select only one type of sorting")
	}
	var newData [][]string
	result := []string{}
	var err error

	//сортируем
	for _, t := range s.data {
		newData = append(newData, strings.Split(t, " "))
	}
	if s.flags.month || s.flags.numeric {
		result, err = sorting(s.flags.coloumn, newData, s.flags.numeric)
		if err != nil {
			return nil, err
		}
	} else if s.flags.numericSuffix {
		result, err = sortSuf(s.flags.coloumn, newData)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = sortColoumn(s.flags.coloumn, newData)
		if err != nil {
			return nil, err
		}
	}

	// делаем реверс
	if s.flags.reverse {
		slices.Reverse(result)
	}

	return result, nil
}

func NewSorter() (*Sorter, error) {
	//парсим флаги
	flags, err := NewFlags()
	if err != nil {
		return nil, err
	}

	// проверям файл
	data, err := readLines(flags.path)
	if err != nil {
		return nil, err
	}

	// проверка на отсортированность
	if flags.checkSort {
		if slicesEqual(data, justSort(data)) {
			fmt.Println("sorted")
		} else {
			fmt.Println("not sorted")
		}
		return nil, nil
	}

	//удаление хвостовых пробелов
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		data[i] = strings.TrimLeft(data[i], " ")
	}

	//проверка на уникальность
	if flags.unique {
		data = makeUnique(data)
	}

	return &Sorter{
		flags: flags,
		data:  data,
	}, nil
}
