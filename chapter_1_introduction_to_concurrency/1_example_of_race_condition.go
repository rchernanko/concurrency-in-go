package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v\n", data)
	}
}

/*

Here, lines 8 and 10 are both trying to access the variable data, but there is no guarantee what order this might happen in.

There are three possible outcomes to running this code:

• Nothing is printed. In this case, line 8 was executed before line 10.
• “the value is 0” is printed. In this case, lines 10 and 11 were executed before line 8.
• “the value is 1” is printed. In this case, line 10 was executed before line 8, but line 8 was executed before line 11.

 */
