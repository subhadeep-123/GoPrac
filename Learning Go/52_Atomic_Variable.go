package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wait sync.WaitGroup
var count int64

func increment(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration((rand.Intn(3))) * time.Millisecond)
		atomic.AddInt64(&count, 1)
		fmt.Println(s, i, "Count -> ", count)
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("Last Count Value ", count)
}
