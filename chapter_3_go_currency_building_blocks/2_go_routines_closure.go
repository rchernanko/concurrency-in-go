package main

import (
	"fmt"
	"sync"
)

/*

Page 40-41

We’ve been using a lot of anonymous functions in our examples to create quick goroutine examples. Let’s shift
our attention to closures. Closures close around the lexical scope they are created in, thereby capturing variables.
If you run a closure in a goroutine, does the closure operate on a copy of these variables, or the
original references? Let’s give it a try and see:

What do you think the value of salutation will be: “hello” or “welcome”? Let’s run it and find out:

welcome

Interesting! It turns out that goroutines execute within the same address space they were created in, and so
our program prints out the word “welcome.”

 */

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()
	fmt.Println(salutation)



}
