package main

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

func ChannelsMain() {

	myChannel := make(chan string)
	anotheChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()
	go func() {
		anotheChannel <- "mango"
	}()

	select {
	case msgFromChannel := <-myChannel:
		fmt.Println(msgFromChannel)
	case msgFromAnotherChannel := <-anotheChannel:
		fmt.Println(msgFromAnotherChannel)

	}
}
