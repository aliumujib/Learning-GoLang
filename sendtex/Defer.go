package main

import (
	"fmt"
)

func doSomething() {
	//defer statements are executed in First In Last Out order
	defer fmt.Println("Done working!!")
	defer fmt.Println("All Done working!!")
	fmt.Println("Doing some stuff, who knows what I am doing?")
}

func main() {
	doSomething()
}
