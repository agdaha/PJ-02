package pipeline

import (
	"bufio"
	"fmt"
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
			fmt.Println("Программа завершила работу!")
			return
		}
		i, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Программа обрабатывает только целые числа!")
			continue
		}
		output <- i
	}
}
