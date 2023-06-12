// To generate a unique ID, you can use a combination of random integers and
// the strconv package to convert the integers to strings:

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate a random 8-digit ID
	id := ""
	for i := 0; i < 8; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}

	fmt.Println(id)
}
