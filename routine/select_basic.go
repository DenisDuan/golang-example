package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
    p :=fmt.Println
	then := time.Now()

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one done"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two done"
	}()

	for i:=0; i<2; i++ {
		select {
			case msg1 := <-c1:
				p("received:", msg1)
			case msg2 := <-c2:
				p("received:", msg2)
		}
	}

	p("Finish time: ", time.Now().Sub(then).Seconds())
}