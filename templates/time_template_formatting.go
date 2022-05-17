package main

import (
	"fmt"
	"time"
)

func main() {
	time_ := time.Now()
	fmt.Println(time_)

	formattedTime := time_.Format(time.Kitchen)
	fmt.Println(formattedTime)

	formattedTime1 := time_.Format(time.RFC850)
	fmt.Println(formattedTime1)
}
