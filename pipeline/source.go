package pipeline

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Source(done chan interface{}, output chan<- int) {
	defer close(done)
	defer close(output)
	scanner := bufio.NewScanner(os.Stdin)
	var data string
	for {
		scanner.Scan()
		data = scanner.Text()
		if strings.EqualFold(data, "exit") {
			//fmt.Println("Программа завершила работу!")
			log.Println("Source: Завершил работу!")
			return
		}
		i, err := strconv.Atoi(data)
		if err != nil {
			//fmt.Println("Программа обрабатывает только целые числа!")
			log.Println("Source: Программа обрабатывает только целые числа!")
			continue
		}
		log.Printf("Source: Передача %v", i)
		output <- i
	}
}
