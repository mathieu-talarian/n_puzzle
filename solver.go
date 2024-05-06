package main

import (
	"fmt"
	"log"
)

var costFunction CostFunction

// Start function init astar
func Start(p *Puzzle, h uint, c uint) {
	a := NewAstar(p, h, c)
	costFunction = FindCostFunction(c)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	if n, err := runN(a); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
		fmt.Println("Number of turns:", a.Turns)
		fmt.Println("Max state:", a.MaxState)
	}
}

const (
	//No action
	No = iota
)

func runN(a *Astar /* , FCost */) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	for a.OpenList.Size() > 0 {
		node := a.OpenList.DeleteMin()
		state := decompute(node.(*Node).State)
		uuid := state.CreateUUID()

		if *node.(*Node).H == 0 {
			return node.(*Node), nil
		}

		(*a).Turns++
		node.(*Node).Execute(a, uuid, state)
		num := a.OpenList.Size()
		if num > int(a.MaxState) {
			a.MaxState = uint(num)
		}
		if a.ClosedList == nil {
			a.ClosedList = NewBst(uuid)
		} else {
			a.ClosedList.Insert(uuid)
		}
		state = nil
		node = nil
	}
	return
}

func move(action Action, state *Puzzle, astar *Astar, n *Node, results chan<- *Node) {
	tile := state.Zero.ToTile(state.Size)
	size := state.Size
	if tile.TestAction(action.Value, size) {
		state.Move(action.Value)
		h, err := astar.HeuristicFunction(state, astar.Goal)
		if err != nil {
			log.Fatal(err)
		}
		results <- NewNode(&action.Name, *n.G+1, uint(h), n, state)
	} else {
		results <- nil
	}
	tile = nil
}

func add(newNode *Node, a *Astar, uuid BstString) {
	if newNode != nil {
		if !newNode.AlreadyClosed(a.ClosedList, uuid) {
			a.OpenList.Insert(newNode)
		} else {
			newNode = nil
		}
	}
}

func worker(id <-chan int, puzzle *Puzzle, a *Astar, n *Node, results chan<- *Node) {
	move(L[<-id], puzzle, a, n, results)
	puzzle = nil
	a = nil
	n = nil
}
