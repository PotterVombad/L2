package pkg

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type Grep struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func NewGrep() (*Grep, error) {
	after := flag.Int("A", 0, "Print N lines after each match")
	before := flag.Int("B", 0, "Print N lines before each match")
	context := flag.Int("C", 0, "Print N lines before and after context")
	count := flag.Bool("c", false, "Count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Invert the sense of matching")
	fixed := flag.Bool("F", false, "Exact match")
	lineNum := flag.Bool("n", false, "Print line numbers")

	// парсинг флагов
	flag.Parse()

	// определение паттерна
	var pattern string
	if flag.NArg() > 0 {
		pattern = flag.Arg(0)
	} else {
		return nil, errors.New("give pattern argument")
	}

	return &Grep{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		pattern:    pattern,
	}, nil
}

func (g *Grep) Run() error{

	// идекс для поиска строк
	lineIndex := 0
	countNum := 0

	// подготовка регулярного выражения в зависимости от флагов
	var re *regexp.Regexp
	if g.fixed {
		g.pattern = regexp.QuoteMeta(g.pattern)
	}
	if g.ignoreCase {
		re = regexp.MustCompile("(?i)" + g.pattern)
	} else {
		re = regexp.MustCompile(g.pattern)
	}

	// чтение входных данных по строкам
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input data: %v", err)
	}

	for _, line := range lines {

		// Проверка совпадения
		matched := re.MatchString(line)

		var flag bool
		// Применение флагов
		if (g.invert && !matched) || (!g.invert && matched) {
			countNum++
			if g.before > 0 || g.context > 0 {
				printContextLines(lines, g.before, g.context, lineIndex, flag)
			}
			printMatchLine(lines, lineIndex, countNum, g.count, g.lineNum)
			if g.before > 0 || g.context > 0 {
				flag = true
				printContextLines(lines, g.after, g.context, lineIndex, flag)
			}
		}

		lineIndex++
	}
	return nil
}

// функция для печати совпавшей строки
func printMatchLine(lines []string, index, countNum int, count, lineNum bool) {
	if count {
		fmt.Printf("count: %d\n", countNum)
	}
	if lineNum {
		index++
		fmt.Printf("index: %d\n", index)
	}
	index--
	fmt.Println(lines[index])
}

// функция для печати контекстных строк
func printContextLines(lines []string, number, context, index int, flag bool) {
	result := 0
	if flag {
		if context+number+index > len(lines)-1 {
			result = len(lines) - 1
		} else {
			result = context + number
		}
		for i := index + 1; i <= result; i++ {
			fmt.Println(lines[i])
		}
	} else {
		if index-context-number < 0 {
			result = 0
		} else {
			result = index - context - number
		}
		for i := result; i < index-1; i++ {
			fmt.Println(lines[i])
		}
	}
}
