// Aceita um conjuto de dados e o reduz a um tamanho fizo menor
// Hashes são usadas em tudos em programação, principalnebte  para buscar DADOS e DETECTAR
//Em GO São dividas em CRIPTOGRAFADAS e NÂO CRIPTOGRAFADAS
// LISTA das não CRIP: adler32. crc32, crc64
// nessse código vamos criar uma Hash e escrever nossos dados nele

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//criar a hash
	h := crc32.NewIEEE()
	//escrever nossos dados no hash
	h.Write([]byte("Código com pacote hash"))
	//calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
