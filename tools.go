package main

import (
	"container/list"
	"fmt"
)

// PrintList prints value of list
func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// PrintAddr print add of interface
func PrintAddr(i ...interface{}) {
	for _, a := range i {
		fmt.Printf("%p ", &a)
	}
	fmt.Println()
}
