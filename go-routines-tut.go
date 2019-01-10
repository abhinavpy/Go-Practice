package main

import "fmt"
import "time"

func compute(value int) {
	for i:=0; i<value;i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Concurrency with go routines\n");
	go compute(5);
	go compute(5);

	fmt.Scanln();
}
