package pipeline

import (
	"fmt"
	"sync"
)

func Consumer(input <-chan int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range input {
			fmt.Printf("Получено: %v \n", data)
		}
	}()
	wg.Wait()
}
