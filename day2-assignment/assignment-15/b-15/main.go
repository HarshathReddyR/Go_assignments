package main

import (
	"fmt"
	"time"
)

func Sender(ch chan int) {
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()
}
func Receiver(ch chan int) {

	go func(ch chan int) {
		for v := range ch {
			fmt.Println(v)
		}
	}(ch)
}

func main() {
	ch := make(chan int, 5)
	Sender(ch)
	Receiver(ch)
	time.Sleep(5 * time.Second)
}
