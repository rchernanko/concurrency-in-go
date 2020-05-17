package main

import (
	"fmt"
	"sync"
	"time"
)

/*

A deadlocked program is one in which all concurrent processes are waiting on one another.
In this state, the program will never recover without outside intervention.

 */

type value struct {
	mu sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock() //1
		defer v1.mu.Unlock() //2

		time.Sleep(2 * time.Second) //3
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value + v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

/*

1 - Here we attempt to enter the critical section for the incoming value.

2 - Here we use the defer statement to exit the critical section before printSum returns.

3 - Here we sleep for a period of time to simulate work (and trigger a deadlock).

If you were to try and run this code, you’d probably see:

	"fatal error: all goroutines are asleep - deadlock!" (I do see this locally)

Why? If you look carefully, you’ll see a timing issue in this code.

Essentially, we have created two gears that cannot turn together: our first call to printSum locks a and then
attempts to lock b, but in the meantime our second call to printSum has locked b and has attempted to lock a.

Both goroutines wait infinitely on each other.

 */