// closere : função "CHAMAR" uma Varável que está em outra função

package main

import "fmt"

func main() {

	x := 0
	numero := func() int {
		x++ // Incrementi soma 1
		return x
	}
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}
