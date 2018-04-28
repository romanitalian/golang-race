package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// $ go run main_waitgroup.go
//1000000

//$ go run -race main_waitgroup.go
//1000000

type atomicCounterWg struct {
	val int64
}

func (c *atomicCounterWg) Add(x int64) {
	atomic.AddInt64(&c.val, x)
	runtime.Gosched()
}

func (c *atomicCounterWg) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounterWg{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second)
	wg.Wait()
	fmt.Println(counter.Value())
}
