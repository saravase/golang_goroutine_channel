package tutorial

import (
	"fmt"
	"time"
)

func publishClose(text string, delay time.Duration) (wait <-chan struct{}) {

	ch := make(chan struct{})

	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch)
	}()
	return ch
}

func Tut02Chan() {
	fmt.Println("START")
	ch := publishClose("12th Result Out", 2*time.Second)
	fmt.Println(<-ch)
	fmt.Println("END")
}
