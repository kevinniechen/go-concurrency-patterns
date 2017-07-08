//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang channels as handles on a service (working in lockstep)

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := generator("Hello")
	ch2 := generator("Bye")
	for i := 0; i < 5; i++ {
		fmt.Println(<- ch1)
		fmt.Println(<- ch2)
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