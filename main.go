package main

import (
	"fmt"
	"time"
)

func ping(channel chan string) {
	for {
		channel <- "ping"
	}
}

func pong(channel chan string) {
	for {
		channel <- "pong"
	}
}

func printMessage(channel chan string) {
	for {
		message := <-channel
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	channel := make(chan string)

	go ping(channel)
	go pong(channel)
	go printMessage(channel)

	fmt.Println("Press enter to exit...")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exiting...")
}