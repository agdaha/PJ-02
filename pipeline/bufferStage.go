package pipeline

import (
	"PJ-02/buffer"
	"time"

	"log"
)

func WriteToBuffer(done <-chan interface{}, input <-chan int, buffer *buffer.RingIntBuffer) {
	for {
		select {
		case data := <-input:
			buffer.Push(data)
			log.Printf("Buffer: Получено %v", data)
		case <-done:
			log.Println("Buffer: writer завершил работу")
			return
		}
	}
}

func ReadFromBuffer(done <-chan interface{}, output chan<- int,
	buffer *buffer.RingIntBuffer, bufferEmptyingInterval time.Duration) {

	defer close(output)
	for {
		select {
		case <-time.After(bufferEmptyingInterval):
			datas := buffer.Get()
			if datas != nil {
				for _, data := range datas {
					select {
					case output <- data:
						log.Printf("Buffer: Передано %v", data)
					case <-done:
						log.Println("Buffer: reader завершил работу")
						return
					}
				}
			}
		case <-done:
			log.Println("Buffer: reader завершил работу")
			return
		}
	}
}
