package main

import (
	"PJ-02/buffer"
	"PJ-02/pipeline"
	"time"
)

const bufferSize = 12

const bufferEmptyingInterval time.Duration = 10 * time.Second

func main() {
	sourceChanel := make(chan int)
	stage1outChanel := make(chan int)
	stage2outChanel := make(chan int)
	outputChanel := make(chan int)
	done := make(chan interface{})

	buffer := buffer.NewRingIntBuffer(bufferSize)

	go pipeline.Source(done, sourceChanel)
	go pipeline.ProcessingStage(done, sourceChanel, stage1outChanel, NegativeFilter)
	go pipeline.ProcessingStage(done, stage1outChanel, stage2outChanel, FilterDivisibility3)
	go pipeline.WriteToBuffer(done, stage2outChanel, buffer)
	go pipeline.ReadFromBuffer(done, outputChanel, buffer, bufferEmptyingInterval)

	pipeline.Consumer(outputChanel)
}
