package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id   int64
	Text string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from thread 1"}
			// time.Sleep(time.Second * 2)
			c1 <- msg
		}
	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from thread 2"}
			// time.Sleep(time.Second * 3)
			c2 <- msg
		}
	}()

	// for i := 0; i < 100; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received from thread 1: id: %d %s\n", msg1.Id, msg1.Text)
		case msg2 := <-c2:
			fmt.Printf("received from thread 2: id: %d %s\n", msg2.Id, msg2.Text)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
	// }
}
