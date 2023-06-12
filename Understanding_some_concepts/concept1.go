// Make a program that rolls a dice (1 to 6)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func random(min int, max int) int {
// 	return rand.Intn(max-min) + min
// }

// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	diceRoll := random(1, 6)
// 	fmt.Printf("Dice roll result: %d\n", diceRoll)
// }

// Can you generate negative numbers?
func RandInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}

func main() {
	randNeg := RandInt(-2, 10)

	fmt.Println(randNeg)
}
