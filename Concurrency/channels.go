package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doworkChannel(d time.Duration, resch chan string) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("work is done!")
	resch <- fmt.Sprintf("result of work %d", rand.Intn(100))
}

func Add(a, b int) int {
	return a + b
}

func main() {
	start := time.Now()
	resultChanel := make(chan string)
	go doworkChannel(time.Second*2, resultChanel)
	go doworkChannel(time.Second*4, resultChanel)
	go doworkChannel(time.Second*6, resultChanel)

	res1 := <-resultChanel
	fmt.Println("res1   ", res1)
	res2 := <-resultChanel
	fmt.Println("res2   ", res2)
	res3 := <-resultChanel
	fmt.Println("res3   ", res3)

	fmt.Printf("work done %v seconds\n", time.Since(start))
}
