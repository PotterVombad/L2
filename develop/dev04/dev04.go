package dev04

import (
	"sort"
	"strings"
)

// реализуем функции для интерфейса sort
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// создаем множество
func makeSet(s []string) []string {
	dict := make(map[string]struct{})
	for _, v := range s {
		_, ok := dict[v]
		if !ok {
			dict[v] = struct{}{}
		}
	}
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return keys
}

// функция поиска анаграм
func FindAnagram(input *[]string) *map[string]*[]string {
	first := make(map[string][]string)
	result := make(map[string]*[]string)
	// добавление в первичную мапу значений по ключу отсортированного слайса рун
	for _, word := range *input {
		word = strings.ToLower(word)
		var runeSlice sortRunes = []rune(word)
		sort.Sort(runeSlice)
		first[string(runeSlice)] = append(first[string(runeSlice)], word)
	}
	// добавление значений в результирующую мапу с учетом всех требований
	for _, value := range first {
		if len(value) > 1 {
			key := value[0]
			valueCp := makeSet(value)
			sort.Strings(valueCp)
			result[key] = &valueCp
		}
	}
	return &result
}
