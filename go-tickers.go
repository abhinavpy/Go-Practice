//Timers are for when you want to do
//something repeatedly at regular intervals
//Here's an example of a ticker that ticks periodically
//until we stop it.
//Tickers use a similar mechanism to timers: a channel 
//that is sent values
//Here we'll use the range builtin on the 
//channel to iterate over the values as they arrive 
//every 500ms.

package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
