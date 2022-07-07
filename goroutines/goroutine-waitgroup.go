package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1) //kaç tane goroutine bekleyeceğimizi bildiriyoruz

	go func() {
		fmt.Println("hello from goroutine")
		wg.Done() //goroutine id done
	}()

	wg.Wait() // blocking

	fmt.Println("hello from main")
}
