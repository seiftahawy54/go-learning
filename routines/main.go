package main

import (
	"fmt"
	"time"
)

func printer(message string, doneChan chan bool) {
	fmt.Println(message)
	doneChan <- true
}

func slowerPrinter(message string, doneChan chan bool) {
	time.Sleep(20 * time.Millisecond)
	fmt.Println(message)
	doneChan <- true
}

func main() {
	doneChan := make([]chan bool, 3)

	doneChan[0] = make(chan bool)
	go printer("This is normal message 1", doneChan[0])
	
	doneChan[1] = make(chan bool)
	go printer("This is normal message 2", doneChan[1])
	
	doneChan[2] = make(chan bool)
	go slowerPrinter("This is slower message 3", doneChan[2])

	for _, done := range doneChan {
		<- done
	}
}