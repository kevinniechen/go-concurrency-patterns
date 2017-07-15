//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang generator pattern: functions that return channels

package main

import (
	"fmt"
	"time"
)

// goroutine is launched inside the called function (more idiomatic)
// multiple instances of the generator may be called

func main() {
	ch := generator("Hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<- ch)
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