package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go squares(ch, 5)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("End.")

}

func squares(ch chan int, num int) {
	for i := 0; i < num; i++ {
		ch <- i * i
		time.Sleep(1 * time.Millisecond)
	}
	close(ch)
}
