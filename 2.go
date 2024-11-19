package main

//это второе задание без дедлока
import (
	"fmt"
	"strconv"
	"sync"
)

var numbers = make(chan int, 10)
var strings = make(chan string, 10)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)

	// Запуск горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range numbers {
				str := strconv.Itoa(num)
				strings <- str
			}
		}()
	}

	// Закрываем канал
	go func() {
		wg.Wait()
		close(strings)
	}()

	//вывод на экран
	for str := range strings {
		fmt.Println(str)
	}
}
