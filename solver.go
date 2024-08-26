package main

import (
	"fmt"
	"log"
)

var costFunction CostFunction

// Start function init astar
func Start(puzzle *Puzzle, heuristic uint, cost uint) {
	astar := NewAstar(puzzle, heuristic, cost)
	costFunction = FindCostFunction(heuristic)
	if !astar.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	if node, err := runN(astar); err != nil {
		log.Fatal(err)
	} else {
		node.PrintResult()
		fmt.Println("Number of turns:", astar.Turns)
		fmt.Println("Max state:", astar.MaxState)
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
}

func add(newNode *Node, a *Astar, uuid BstString) {
	if newNode != nil {
		if !newNode.AlreadyClosed(a.ClosedList, uuid) {
			a.OpenList.Insert(newNode)
		}
	}
}

func worker(id <-chan int, puzzle *Puzzle, a *Astar, n *Node, results chan<- *Node) {
	move(L[<-id], puzzle, a, n, results)
}
