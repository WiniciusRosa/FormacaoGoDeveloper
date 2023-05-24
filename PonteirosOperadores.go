// Ponteiro (ptr) armazenar um valor na memória , mas não é o valor propriamente escrito

package main

import "fmt"

func Inicial(xPtr *int) {

	*xPtr = 0

}

func main() {
	x := 5

	Inicial(&x)
	fmt.Println(x)

}
