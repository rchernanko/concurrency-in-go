package main

/*

Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the
state of the program forward.

Have you ever been in a hallway walking toward another person?
She moves to one side to let you pass, but you’ve just done the same.
So you move to the other side, but she’s also done the same.
Imagine this going on forever, and you understand livelocks.


Actually not going to put the example of the livelock in here...just check out the book

In my opinion, livelocks are more difficult to spot than deadlocks simply because it can appear as if the program
is doing work. If a livelocked program were running on your machine and you took a look at the CPU utilization to
determine if it was doing anything, you might think it was. Depending on the livelock, it might even be emitting
other signals that would make you think it was doing work.
And yet all the while, your program would be playing an eternal game of hallway-shuffle.

 */
