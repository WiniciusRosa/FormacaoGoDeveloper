// panic: erro do progamador e erro de execução do tempo

package main

import "fmt"

func main() {

	defer func() {

		x := recover()

		fmt.Println(x)
	}()
	panic("Panico")
}
