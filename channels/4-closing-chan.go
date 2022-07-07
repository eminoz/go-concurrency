package main

import "fmt"

type name struct {
	username string
	index    int
}

func producer(chnnl chan name) {
	user := []string{"emin", "burak", "ali"}
	for i, i2 := range user {

		chnnl <- name{username: i2, index: i}
	}
	close(chnnl) //bunu sadece göndericilerde kullan
}
func main() {

	ch := make(chan name)
	go producer(ch)
	for {
		i, ok := <-ch
		if ok == false { // chanda veri bitince ok false döner
			break
		}
		fmt.Println(i)

	}
}
