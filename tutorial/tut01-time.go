package tutorial

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

func Tut01Time() {
	Publish("10th Result Out", 5*time.Second)
	fmt.Println("START")

	time.Sleep(10 * time.Second)
	fmt.Println("END")
}
