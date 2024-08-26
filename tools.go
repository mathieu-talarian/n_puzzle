package main

import (
	"container/list"
	"fmt"
)

/* PrintList prints the values of a list */
func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

/* PrintAddr prints the address of the provided interfaces */
func PrintAddr(i ...interface{}) {
	for _, a := range i {
		fmt.Printf("%p ", &a)
	}
	fmt.Println()
}
