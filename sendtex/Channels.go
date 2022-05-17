package main

import "fmt"

func sendThroughChannel(channel chan int, value int) {
	channel <- value * 5
}

func main() {

	channel := make(chan int)

	go sendThroughChannel(channel, 5)
	go sendThroughChannel(channel, 7)

	channelVal1 := <-channel
	channelVal2 := <-channel

	// fancier way
	// 	channelVal1, channelVal2 := <-channel, <-channel

	fmt.Println(channelVal1, channelVal2)
}
