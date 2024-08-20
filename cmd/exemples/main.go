package main

import (
	"fmt"
	"time"
)

func worker(workId int, data chan int){
	for x := range data {
		fmt.Printf("worker %d got %d\n", workId, x)
		time.Sleep(time.Second)
	}

}

func main() {
	ch := make(chan int)

	go worker(1, ch)
	go worker(2, ch)
	for i := range 10 {
		ch <- i
	}

}
