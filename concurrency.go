package main

import (
	"fmt"
	"time"
)

// calculateSum calculates the sum of numbers from 1 to n.
// It runs concurrently and once the sum is calculated, it sends the result to sumChannel.
// The use of channels allows us to synchronize concurrent processes and communicate between them,
// thereby solving potential concurrency issues like data races.
func calculateSum(n int, sumChannel chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	sumChannel <- sum // send the sum to the channel
}

// printNumbers prints numbers from 1 to n with a delay.
// This function runs concurrently and indicates its completion by sending a message to doneChannel.
// The use of channels ensures that the main function can wait for this concurrent process to complete before proceeding.
func printNumbers(n int, doneChannel chan bool) {
	for i := 1; i <= n; i++ {
		fmt.Println("Number:", i)
		time.Sleep(100 * time.Millisecond) // just to simulate some work
	}
	doneChannel <- true // send a message to the channel indicating we're done printing
}

func main() {
	n := 10

	// Creating channels to handle synchronization and communication between concurrent processes.
	sumChannel := make(chan int)
	doneChannel := make(chan bool)

	// Starting goroutines. The "go" keyword allows these functions to run concurrently.
	go calculateSum(n, sumChannel)
	go printNumbers(n, doneChannel)

	// By reading from the channels, we are essentially waiting for the concurrent processes to complete.
	// This helps in ensuring that the program doesn't exit prematurely and that all tasks have completed before we move on.
	sum := <-sumChannel
	done := <-doneChannel

	if done {
		fmt.Printf("The sum of numbers from 1 to %d is %d\n", n, sum)
	}
}
