package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const max int = 10

var wait sync.WaitGroup
var count int

func increment(str string) {
	for i := 0; i < max; i++ {
		value := count
		value++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = value
		fmt.Println(str, i, "Count: ", count)
	}

	wait.Done()
}

func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value", count)
}
