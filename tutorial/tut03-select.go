package tutorial

import (
	"fmt"
)

func Tut03Select() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
		ch2 <- 2
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("Channel 1 Received data")
		case <-ch2:
			fmt.Println("Channel 2 Received data")
		default:
			fmt.Println("No Channel Exist")
			break
		}
	}
}
