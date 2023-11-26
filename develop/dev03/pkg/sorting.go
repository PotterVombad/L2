package pkg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var months = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func sorting(coloumn int, lines [][]string, numeric bool) ([]string, error) {
	var keys []int
	lineMap := map[int][]string{}
	var err error

	for _, line := range lines {
		if len(line)-1 < coloumn {
			return nil, fmt.Errorf("invalid coloumn")
		}
		var value int
		if numeric {
			value, err = strconv.Atoi(line[coloumn])
			if err != nil {
				return nil, fmt.Errorf("invalid coloumn")
			} 
		} else {
			_, ok := months[line[coloumn]]
			if !ok {
				return nil, fmt.Errorf("invalid coloumn")
			}
			value = months[line[coloumn]]
		} 
		keys = append(keys, value)
		lineMap[value] = line
	}

	sort.Ints(keys)
	var result []string

	for _, key := range keys {
		res := strings.Join(lineMap[key], " ")
		result = append(result, res)
	}

	return result, nil
}

func sortSuf(coloumn int, lines [][]string) ([]string, error) {
	var keys []float64
	lineMap := map[float64][]string{}
	var err error

	for _, line := range lines {
		if len(line)-1 < coloumn {
			return nil, fmt.Errorf("invalid coloumn")
		}
		var value float64
		value, err = strconv.ParseFloat(line[coloumn], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid coloumn")
		} 
		keys = append(keys, value)
		lineMap[value] = line
	}

	sort.Float64s(keys)
	var result []string

	for _, key := range keys {
		res := strings.Join(lineMap[key], " ")
		result = append(result, res)
	}

	return result, nil
}

func sortColoumn(coloumn int, lines [][]string) ([]string, error) {
	var keys []string
	lineMap := map[string][]string{}

	for _, line := range lines {
		if len(line)-1 < coloumn {
			return nil, fmt.Errorf("invalid coloumn")
		}
		keys = append(keys, line[coloumn])
		lineMap[line[coloumn]] = line
	}

	sort.Strings(keys)
	var result []string

	for _, key := range keys {
		res := strings.Join(lineMap[key], " ")
		result = append(result, res)
	}

	return result, nil
}
