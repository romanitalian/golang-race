package main

import (
	"time"
	"fmt"
)

type intCounter int64

func (c *intCounter) Add(x int64) {
	*c++
}
func (c *intCounter) Value() int64 {
	return int64(*c)
}
func main() {
	counter := intCounter(0)
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(counter.Value())
}

