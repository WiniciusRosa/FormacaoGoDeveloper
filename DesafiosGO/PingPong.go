package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Println("BateuVoltou")
		} else if i%3 == 0 {
			fmt.Println("Bateu")
		} else if i%5 == 0 {
			fmt.Println("Voltou")
		} else {
			fmt.Println(i)

		}

	}
}
