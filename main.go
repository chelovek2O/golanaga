package main

//дедлок
import (
	"fmt"
	"strconv"
	"sync"
)

var (
	numbers = make(chan int, 10)
	strings = make(chan string, 10)
)

func main() {
	var wg sync.WaitGroup

	// Запись чисел
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
				strings <- str // Здесь происходит потенциальная блокировка
			}
		}()
	}

	go func() {
		wg.Wait()
		close(strings)
	}()

	//вывод на экран
	for str := range strings {
		fmt.Println(str)
	}
}

//дедлок происходит невсегда я навсякий случай приложил картинку как он появился
