package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"sync"
)

//Note the below is not idiomatic go but very simply demonstrates memory access synchronisation

func main() {
	var memoryAccess sync.Mutex //1
	var value int
	go func() {
		memoryAccess.Lock() //2
		value++
		memoryAccess.Unlock() //3
	}()

	memoryAccess.Lock() //4
	if value == 0 {
		fmt.Printf("the value is %v\n", value)
	} else {
		fmt.Printf("the value is %v\n", value)
	}
	memoryAccess.Unlock() //5
}

/*

1 - Here we add a variable that will allow our code to synchronize access to the data variable’s memory.

2 - Here we declare that until we declare otherwise, our goroutine should have exclusive access to this memory.

3 - Here we declare that the goroutine is done with this memory.

4 - Here we once again declare that the following conditional statements should have exclusive access to the data variable’s memory.

5 - Here we declare we’re once again done with this memory.

NOTE - You may have noticed that while we have solved our data race, we haven’t actually solved our race condition!
The order of operations in this program is still non-deterministic

Synchronizing access to the memory in this manner also has performance ramifactions

*/
