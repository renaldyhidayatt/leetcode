package main

import "fmt"

func PushChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)
}

func ForRangeChannel() {
	ch := make(chan string)
	go PushChannel(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
