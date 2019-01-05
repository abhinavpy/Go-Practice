package main

import ("fmt"
		)

type car struct {
	gas_pedal uint64 //min 0, max 65535
	brake_pedal uint16 
	steering_whell int16
	top_speed float64
}

func main() {
	a_car := car{gas_pedal: 22341,
				 brake_pedal: 0, 
				 steering_whell: 12561, 
				 top_speed: 225.0}
	fmt.Println(a_car.gas_pedal)
}