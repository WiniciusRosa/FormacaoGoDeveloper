package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {

	var c float64 = float64(a)

	fmt.Printf("O valor de C é: %g (%T) e  o valor de b é: %s (%T)", c, c, b, b)

	//fmt.Printf("x= %d, %T  e  y= %g, %T", x, x, y, y)
}
