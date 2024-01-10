package main

import (
	"fmt"
	"sync"
	"time"
)

func dowork(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println(d)
	fmt.Println("work is done!")
	wg.Done()
}

func WaitMain() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go dowork(time.Second*1, &wg)
	go dowork(time.Second*1, &wg)
	wg.Wait()
	fmt.Printf("work done %v seconds\n", time.Since(start))
}
