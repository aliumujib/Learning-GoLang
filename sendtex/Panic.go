package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup //like a CountDownLatch or Sephamore .. from Java

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("New Panic occured", r)
	}
	waitGroup.Done()
}

func say(text string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		if i == 2 {
			panic("Oh no, number 2 showed up")
		}
		fmt.Println(text)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	waitGroup.Add(1)
	go say("Sup!")
	waitGroup.Add(1)
	go say("How you doin!")
	waitGroup.Wait()
}
