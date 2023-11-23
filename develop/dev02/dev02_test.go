package dev02

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	task   string
	result string
	err    error
}

func TestUnpacked(t *testing.T) {
	tables := []test{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("некорректная строка")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}
	for _, table := range tables {
		result, err := Unpacke(table.task)

		assert.Equal(t, table.result, result)
		assert.Equal(t, table.err, err)
	}
}
