package dev07

import (
	"sync"
)

func or(channels ...<-chan interface{}) chan interface{} {
	result := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			defer wg.Done()
			for {
				_, ok := <-ch
				if !ok {
					break
				}
			}
		}(ch)
	}
	wg.Wait()
	close(result)
	return result
}
