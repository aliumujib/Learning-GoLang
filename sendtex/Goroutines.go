package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup //like a CountDownLatch or Sephamore .. from Java

func say(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 100)
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)
	go say("Yo!")
	waitGroup.Add(1)
	go say("Sup!")
	waitGroup.Add(1)
	go say("Show!")

	waitGroup.Wait()
}
