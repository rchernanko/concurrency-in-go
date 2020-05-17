package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v\n", data)
	}
}

/*

Have we solved our data race? No. In fact, it’s still possible for all three outcomes to arise from this program, just increasingly unlikely.

The longer we sleep in between invoking our goroutine and checking the value of data, the closer our program gets to
achieving correctness—but this probability asymptotically approaches logical correctness; it will never be logically correct.

In addition to this, we’ve now introduced an inefficiency into our algorithm. We now have to sleep for one second to
make it more likely we won’t see our data race. If we utilized the correct tools, we might not have to wait at all,
or the wait could be only a microsecond.

The takeaway here is that you should always target logical correctness.
Introducing sleeps into your code can be a handy way to debug concurrent programs, but they are not a solution.

 */