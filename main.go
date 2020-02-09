package main

import (
	"fmt"
	"sync"
	"strconv"
	"github.com/opesun/goquery"
	"strings"
	"time"
)

const (
	PAGES int = 2 //кол-во "воркеров"
	BASE_URL string = "https://www.avito.ru/rossiya/avtomobili/audi?cd=1&p="
)

func main() {
	c := make(chan string)
	go grab(c)

	for s := range c {
		fmt.Println("Начинаем читать:  \n")
		fmt.Println(s, "\n")
	}
}

func grab(c chan string) {
	var wg sync.WaitGroup
	for i := 1; i <= PAGES; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			var url string = BASE_URL + strconv.Itoa(page)

			fmt.Println("Проверка: page = ", page, " URL = ", url, "\n")

			x, err := goquery.ParseUrl(url)
			if err != nil {
				fmt.Println("Не читается: page = ", page, "URL = ", url, "\n")
				return
			}

			for i := 0; i < 1; i++ {
				if s := strings.TrimSpace(x.Find(".snippet-title a").Text()); s != "" {
					c <- s
					//c <- "123\n"
				}
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Запущено воркеров: ", PAGES)

	wg.Wait()
	close(c)
}
