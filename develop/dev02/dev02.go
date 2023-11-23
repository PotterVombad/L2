package dev02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpacke(s string) (string, error) {
	_, err := strconv.Atoi(string(s))
	if err == nil {
		return "", errors.New("некорректная строка")
	}
	var answer strings.Builder
	var escape bool
	var last rune
	input := []rune(s)
	for _, val := range input {
		if unicode.IsDigit(val) && !escape {
			num, err := strconv.Atoi(string(val))
			if err != nil {
				return "", err
			}
			repeat := strings.Repeat(string(last), num-1)
			answer.WriteString(repeat)
		} else {
			escape = val == '\\' && !escape
			if !escape {
				answer.WriteRune(val)
			}
		}
		last = val
	}
	return answer.String(), nil
}
