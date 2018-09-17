//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang channels

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go channel_print("Hello", ch)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // ends of channel block until both are ready
						  // NOTE: golang supports buffered channels, like mailboxes (no sync)
	}
	fmt.Println("Done!")
}

func channel_print(msg string, ch chan<- string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}
