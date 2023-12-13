package pipeline

import (
	"log"
	"reflect"
	"runtime"
)

func ProcessingStage(done <-chan interface{}, input <-chan int, output chan<- int, process func(int) (int, bool)) {
	defer close(output)
	funcName := runtime.FuncForPC(reflect.ValueOf(process).Pointer()).Name()
	for {
		select {
		case inData := <-input:
			outData, ok := process(inData)
			if ok {
				output <- outData
				log.Printf("ProcessingStage(%v): Передано %v", funcName, outData)
			} else {
				log.Printf("ProcessingStage(%v): Отсеяно %v", funcName, inData)
			}

		case <-done:
			log.Println("ProcessingStage: Завершение")
			return
		}
	}
}
