//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//  Daisy chaining goroutines... all routines at once, so if one is fulfilled
//  everything after should also have their blocking commitment (input) fulfilled

package main

import (
	"fmt"
)

// takes two int channels, stores right val (+1) into left
func f(left, right chan int) {
	left <- 1 + <-right // bafter 1st right read, locks until left read
}

func main() {
	const n = 10000

	// construct an array of n+1 int channels
	var channels [n + 1]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}

	// wire n goroutines in a chain
	for i := 0; i < n; i++ {
		go f(channels[i], channels[i+1])
	}

	// insert a value into right-hand end
	go func(c chan<- int) { c <- 1 }(channels[n])

	// get value from the left-hand end
	fmt.Println(<-channels[0])
}
