package main

import (
	"fmt"
	"time"
)

func main() {
	// Define and execute an outer anonymous function as a Goroutine
	go func() {
		fmt.Println("Outer function started.")
		for i := 1; i <= 3; i++ {
			fmt.Println("Inner:", i)
			time.Sleep(200 * time.Millisecond)
		}
		// Define and execute an inner anonymous function
		go func() {
			fmt.Println("Inner function started.")
			for i := 1; i <= 3; i++ {
				fmt.Println("Inner:", i)
				time.Sleep(200 * time.Millisecond)
			}
			fmt.Println("Inner function finished.")
		}()

		fmt.Println("Outer function finished.")
	}()

	// Allow some time for the Goroutines to execute
	time.Sleep(1 * time.Second)

	fmt.Println("Main function finished executing.")
}
