package main

import (
	"fmt"
	"sync"
	"time"
)

// Function that runs as a Goroutine
func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when done
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500) // Simulating some work
	}
}

// Function to demonstrate Goroutines with Channels
func sendData(ch chan string) {
	time.Sleep(time.Second * 1) // Simulate delay
	ch <- "Hello from Goroutine using Channel!"
}

func main() {
	fmt.Println("Main function started")

	// Using sync.WaitGroup to wait for Goroutines to finish
	var wg sync.WaitGroup

	// Starting multiple Goroutines
	wg.Add(2) // Increment WaitGroup counter by 2
	go printMessage("Goroutine 1", &wg)
	go printMessage("Goroutine 2", &wg)

	// Creating a Channel for communication
	messageChannel := make(chan string)

	// Starting a Goroutine that sends data to the channel
	go sendData(messageChannel)

	// Receiving data from the channel
	msg := <-messageChannel
	fmt.Println(msg) // Prints "Hello from Goroutine using Channel!"

	// Wait for all Goroutines to complete
	wg.Wait()

	fmt.Println("Main function finished")
}
