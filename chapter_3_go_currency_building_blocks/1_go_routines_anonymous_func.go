package main

import (
	"fmt"
	"sync"
)

/*

Page 39-40

Go follows a model of concurrency called the fork-join model. The word fork refers to the fact that at any point
in the program, it can split off a child branch of execution to be run concurrently with its parent.
The word join refers to the fact that at some point in the future, these concurrent branches of execution will
join back together. Where the child rejoins the parent is called a join point.

Nice graphic on page 39 of the book.

This example will deterministically block the main goroutine until the goroutine hosting the sayHello function terminates.

 */

func main() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)

	go sayHello()

	wg.Wait() //this is the join point
}
