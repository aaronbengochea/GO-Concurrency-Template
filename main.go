package main

import (
	"fmt"
	"sync"
)

// Resource:
// https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d \n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	// calls done at function completion, reducing the count on Add().
	defer wg.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d \n", idx)
	}
}

func generateNumbersChannel(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("sending %d to channel \n", idx)
		ch <- idx
	}
}

func printNumbersChannel(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("read %d from channel \n", num)
	}
}

func main() {
	// Declare waitgroup
	var wg sync.WaitGroup
	sampleChannel := make(chan int)

	// Instructs the waitgroup to wait for 2 Done calls before considering the
	// group finished. Best practice to always let waitgroup know how many goroutines
	// to expect.
	wg.Add(2)

	// These represent two goroutines which are ran concurrently
	//go printNumbers(&wg)
	//go generateNumbers(3, &wg)

	fmt.Printf("Waiting for goRoutines to finish.. \n")

	// Waits for goRoutines waitgroup to finish running before continuing
	wg.Wait()

	fmt.Printf("Done!")

}
