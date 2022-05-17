package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func sendThroughChannel(channel chan int, value int) {
	defer waitGroup.Done()
	channel <- value * 5
}

func main() {

	channel := make(chan int, 11) // the 11 is a buffer, read more at https://www.geeksforgeeks.org/buffered-channel-in-golang/#:~:text=Buffered%20channels%20allows%20to%20accept,the%20buffer%20will%20be%20empty.

	for number := 0; number < 11; number++ {
		waitGroup.Add(1)
		go sendThroughChannel(channel, number)
	}

	waitGroup.Wait()
	close(channel)

	for result := range channel {
		fmt.Println(result)
	}

}
