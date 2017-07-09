//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//  In non deterministic select control block, 1 second timer (created each iteration) may
//  time out if channel does not return a string in a second

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hi!")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // time.After returns a channel that waits N time to send a message
			fmt.Println("Waited too long!")
			return
		}
	}
}

func generator(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
