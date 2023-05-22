package main

import "fmt"

var x int
var y float64

func main() {

	x = 10
	y = 10.9
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Printf("x= %d, %T  e  y= %g, %T", x, x, y, y)
}
