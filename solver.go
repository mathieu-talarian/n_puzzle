package main

import (
	"fmt"
	"log"
)

var costFunction CostFunction

// Start function init astar
func Start(puzzle *Puzzle, heuristic uint) {
	astar := NewAStarSolver(puzzle, heuristic)
	costFunction = FindCostFunction(heuristic)
	if !astar.IsSolvable() {
		log.Fatal("This puzzle is unsolvable")
	}
	if node, err := runN(astar); err != nil {
		log.Fatal(err)
	} else {
		node.PrintResult()
		fmt.Println("Number of turns:", astar.NumberOfTurns)
		fmt.Println("Max state:", astar.MaxStatesReached)
	}
}

func runN(a *AStarSolver) (q *Node, err error) {
	if err = a.InitializeRootNode(); err != nil {
		return
	}
	for a.OpenNodesHeap.Size() > 0 {
		node := a.OpenNodesHeap.DeleteMin()
		state := decompute(node.(*Node).State)
		uuid := state.CreateUUID()

		if *node.(*Node).H == 0 {
			return node.(*Node), nil
		}

		(*a).NumberOfTurns++
		node.(*Node).Execute(a, uuid, state)
		num := a.OpenNodesHeap.Size()
		if num > int(a.MaxStatesReached) {
			a.MaxStatesReached = uint(num)
		}
		if a.ClosedNodesTree == nil {
			a.ClosedNodesTree = NewBst(uuid)
		} else {
			a.ClosedNodesTree.Insert(uuid)
		}
	}
	return
}

func move(action Action, state *Puzzle, aStar *AStarSolver, n *Node, results chan<- *Node) {
	tile := state.Zero.ToTile(state.Size)
	size := state.Size
	if tile.TestAction(action.Value, size) {
		state.Move(action.Value)
		h, err := aStar.HeuristicFunc(state, aStar.GoalPuzzle)
		if err != nil {
			log.Fatal(err)
		}
		results <- NewNode(&action.Name, *n.G+1, uint(h), n, state)
	} else {
		results <- nil
	}
}

func add(newNode *Node, aStar *AStarSolver, uuid TreeString) {
	if newNode != nil {
		if !newNode.AlreadyClosed(aStar.ClosedNodesTree, uuid) {
			aStar.OpenNodesHeap.Insert(newNode)
		}
	}
}

func worker(id <-chan int, puzzle *Puzzle, aStar *AStarSolver, n *Node, results chan<- *Node) {
	move(ActionsList[<-id], puzzle, aStar, n, results)
}
