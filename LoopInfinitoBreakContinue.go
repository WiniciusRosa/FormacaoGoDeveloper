package main

import "fmt"

func main() {

	for x := 0; x < 20; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)
	}

	//	x := 0
	//
	//	for {
	//		if x < 20 {
	//			x++
	//			fmt.Println("i < 20")
	//		} else {
	//			break
	//		}
	//
	//	}

}
