package dev07

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestOr(t *testing.T) {
	start := time.Now()

	<-or(
		sig(4*time.Second),
		sig(3*time.Second),
		sig(2*time.Second),
		sig(7*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))

	assert.Equal(t, 7, int(time.Since(start).Seconds()))
}
