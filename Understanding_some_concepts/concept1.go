// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func responseSize(url string) {
// 	fmt.Println("Step1: ", url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Step2: ", url)
// 	defer response.Body.Close()

// 	fmt.Println("Step3: ", url)
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Step4: ", len(body))
// }

// func main() {
// 	go responseSize("https://www.golangprograms.com")
// 	go responseSize("https://coderwall.com")
// 	go responseSize("https://stackoverflow.com")
// 	time.Sleep(10 * time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func goroutine1(str string) {
// 	for w := 0; w < 6; w++ {
// 		fmt.Println(str)
// 	}
// }

// func main() {

//		//Calling goroutine
//		go goroutine1("It is a goroutine!!")
//		time.Sleep(1 * time.Second)
//		//Calling normal function
//		goroutine1("It is a function.")
//	}
package main

import "fmt"

// Prints a greeting message using values received in
// the channel
func greet(c chan string) {

	name := <-c // receiving value from channel
	fmt.Println("Hello", name)
}

func main() {

	// Making a channel of value type string
	c := make(chan string)

	// Starting a concurrent goroutine
	go greet(c)

	// Sending values to the channel c
	c <- "World"

	// Closing channel
	close(c)
}
