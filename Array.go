package main

import "fmt"

func main() {

	var x [6]float64
	x[0] = 5.3
	x[1] = 3.3
	x[2] = 7.3
	x[3] = 9.3
	x[4] = 8.3
	x[5] = 10

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

//-----------------------------------------------------//	
	var x [6]float64

	// x[0] = 5.3
	// x[1] = 3.3
	// x[2] = 7.3
	// x[3] = 9.3
	// x[4] = 8.3
	// x[5] = 10
	//
	// var total float64 = 0
	//
	//	for i := 0; i < 6; i++ {
	//		total += x[i]
	//	}
	//
	// fmt.Println(total / 6)
}
