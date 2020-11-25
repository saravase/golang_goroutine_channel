package tutorial

import "fmt"

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
}

func Tut04Race() {
	sharingIsCaring()
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait
	fmt.Println(n) // Output: <unspecified>
}
