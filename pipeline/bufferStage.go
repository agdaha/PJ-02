package pipeline

import (
	"PJ-02/buffer"
	"time"
)

func WriteToBuffer(done <-chan interface{}, input <-chan int, buffer *buffer.RingIntBuffer) {
	for {
		select {
		case data := <-input:
			buffer.Push(data)
		case <-done:
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
					case <-done:
						return
					}
				}
			}
		case <-done:
			return
		}
	}
}
