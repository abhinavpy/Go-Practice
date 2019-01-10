//A go routine is a light weight thread of execution
//Suppose we have a function call f(s). Here's how we'd call 
//that in the usual way, running it synchronously
//To invoke this function in a goroutine,use f(s). This new goroutine 
//will execute concurrently with the calling one.
//You can also start a goroutine for an anonymous function call.
//our two function calls are running asynchronously in separate goroutines 
//now, so execution falls through to here. This Scanln requires we
//press a key before the program exits.
//when we run this program, we see the output of the 
//blocking first, then the interleaved output of the
//two goroutines. This interleaving reflects the goroutines 
//being run concurrently by the go runtime.

package main

import "fmt"

func f(from string) {
	for i:=0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}