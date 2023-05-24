package main

import (
	"container/list"
	"fmt"
)

func main() {
	// CRIAÇÂO
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	//LISTA
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
