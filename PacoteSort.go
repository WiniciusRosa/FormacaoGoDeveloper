package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}
type ParaNome []Dados

func (ps ParaNome) len() int {
	return len(ps)
}
func (ps ParaNome) Less(i, j int) bool { // item na possição i é menor que o item na posição j
	return ps[i].Nome < ps[j].Nome
}
func (ps ParaNome) Swap(i, j int) { //swp: trocar os itens
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	criancas := []Dados{
		{"Juliano", 6},
		{"Winicius", 22},
	}

	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
