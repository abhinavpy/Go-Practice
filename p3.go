package main

import ("fmt")

func add(x,y float64) float64{
	return x+y
}

func main(){
	x,y := 5.5, 6.4
	fmt.Println("the addition of %g and %g is %g", x,y,add(x,y))
}
