package main

import "fmt"

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})

	cancel := func() {
		close(done)
	}

	go func() {
		ch <- 2
		close(ch)
	}()

	return ch, cancel
}

func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		fmt.Println(i)
	}
	cancel()
}
