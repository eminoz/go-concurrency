package main

import (
	"fmt"
	"time"
)

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//reader goroutine
	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	//writer goroutine
	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)
	time.Sleep(time.Second)
	fmt.Println("hello channels")

	/*
		Output is non-deterministic. Scheduler probably will not have time to schedule goroutines.
		So we will not see channel value in the output
	*/
}
