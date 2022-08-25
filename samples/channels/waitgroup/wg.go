package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Doing work")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Doing other work")
	}()

	fmt.Println("Waiting for wait group to complete")
	wg.Wait()
	fmt.Println("Wait group completed")
}