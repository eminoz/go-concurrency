package main

import "fmt"

/*
 	channel senkronize edilmiş bir arraydir
	channel aslında arkada locklanan bir array
*/

func main() {
	//channel initialization
	unbufferedChan := make(chan int) //channellar make ile tanımlanır ve içerisinde taşıyacağı datanın tipi belirlenir
	//unbuffered chan bir blocking işlemidir

	//only can read from channel
	go func(unbufChan <-chan int) {
		//blocks until data arrives
		value := <-unbufChan
		fmt.Println(value)

		//unbufChan <- 5 this is not work here
	}(unbufferedChan)

	unbufferedChan <- 3
}
