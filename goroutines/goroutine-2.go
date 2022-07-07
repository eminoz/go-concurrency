package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 3)
}
