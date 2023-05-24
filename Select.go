// Select funciona como um Switch para canais
// Select permite que você aguarda  operações de vários canais
// Combinar goroutines e canais como select é um recurso poderoso do GO
// Switch
// Para nosso exemplo, selecionaremos em dois canais

package main

// Cada canal receberá um valor após algum tempo, para simular, por exemplo
// o bloqueio de operações RPC em execução em gorountines concorrentes

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "um"

	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}

	}
}
