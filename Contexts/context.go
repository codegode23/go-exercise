package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	//fmt.Println("Doing something!")
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	//ctx := context.TODO() create an empty (or starting) context.
	//You can use this as a placeholder when you’re not sure which context to use.
	ctx := context.Background() //creates an empty context like context.TODO does,
	//but it’s designed to be used where you intend to start a known context.

	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
