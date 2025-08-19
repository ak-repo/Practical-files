package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Select in channels")

	ch1 := make(chan int)

	ch2 := make(chan int)

	done := make(chan bool)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()
	time.Sleep(time.Millisecond * 100)

	t := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case m := <-ch1:
				fmt.Println(m)
			case m := <-ch2:
				fmt.Println(m)

			case m := <-t.C:
				fmt.Println(m)
			case <-done:
				fmt.Println("done")
			}

		}
	}()

	time.Sleep(time.Second * 3)
	done <- true

}
