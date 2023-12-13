package pipeline

func ProcessingStage(done <-chan interface{}, input <-chan int, output chan<- int, procces func(int) (int, bool)) {
	defer close(output)
	for {
		select {
		case inData := <-input:
			outData, ok := procces(inData)
			if ok {
				output <- outData
			}
		case <-done:
			return
		}
	}
}
