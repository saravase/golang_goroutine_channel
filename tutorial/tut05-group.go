package tutorial

import (
	"fmt"
	"sync"
)

func Tut05Group() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Go-routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Go-routine 2")
		wg.Done()
	}()
	wg.Wait()
}
