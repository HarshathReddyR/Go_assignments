package main

import (
	"fmt"
	"sync"
)

func Sender(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	if i == 2 {
		close(ch)
	}
}
func Receiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}

}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 5)

	for i := 0; i <= 2; i++ {
		wg.Add(1)
		// defer wg.Done()
		go Sender(ch, &wg, i)
	}
	for i := 0; i <= 1; i++ {
		wg.Add(1)
		// defer wg.Done()
		go Receiver(ch, &wg)
	}
	// time.Sleep(5 * time.Second)
	wg.Wait()
}
