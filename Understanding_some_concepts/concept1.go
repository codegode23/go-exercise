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

// 	//Calling goroutine
// 	go goroutine1("It is a goroutine!!")
// 	time.Sleep(1 * time.Second)
// 	//Calling normal function
// 	goroutine1("It is a function.")
// }

package main

import (
	"fmt"
	"time"
)

func number() { //creating a goroutine
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond) //making the goroutine sleep for 400 milliseconds
		fmt.Printf("%d ", i)
	}
}

func alphabet() { //creating a goroutine
	for i := 'a'; i <= 'g'; i++ {
		time.Sleep(500 * time.Millisecond) //making the goroutine sleep for 500 milliseconds
		fmt.Printf("%c ", i)
	}
}

func main() {
	go number()                         //calling a goroutine
	go alphabet()                       //calling a goroutine
	time.Sleep(9000 * time.Millisecond) //making the main goroutine sleep for 9000 milliseconds
	fmt.Println("\nProgram terminated!!")
}
