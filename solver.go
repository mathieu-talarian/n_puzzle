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
		node.PrintSolutionPath()
		fmt.Println("Number of turns:", astar.NumberOfTurns)
		fmt.Println("Max state:", astar.MaxStatesReached)
	}
}

func runN(a *AStarSolver) (q *SearchNode, err error) {
	if err = a.InitializeRootNode(); err != nil {
		return
	}
	for a.OpenNodesHeap.Size() > 0 {
		node := a.OpenNodesHeap.DeleteMin()
		state := Decompute(node.(*SearchNode).EncodedState)
		uuid := state.CreateUUID()

		if *node.(*SearchNode).HeuristicCost == 0 {
			return node.(*SearchNode), nil
		}

		(*a).NumberOfTurns++
		node.(*SearchNode).ExecuteSolver(a, uuid, state)
		num := a.OpenNodesHeap.Size()
		if num > int(a.MaxStatesReached) {
			a.MaxStatesReached = uint(num)
		}
		a.ClosedNodesMap[uuid] = struct{}{}
	}
	return
}

func move(action Action, state *Puzzle, aStar *AStarSolver, searchNode *SearchNode, results chan<- *SearchNode) {
	tile := state.ZeroPosition.ToTile(state.Size)
	size := state.Size
	if tile.TestAction(action.Value, size) {
		state.Move(action.Value)
		h, err := aStar.HeuristicFunc(state, aStar.GoalPuzzle)
		if err != nil {
			log.Fatal(err)
		}
		results <- NewSearchNode(&action.Name, *searchNode.HeuristicCost+1, uint(h), searchNode, state)
	} else {
		results <- nil
	}
}

func add(newNode *SearchNode, aStar *AStarSolver, uuid TreeString) {
	if newNode != nil {
		if !newNode.IsAlreadyClosed(aStar.ClosedNodesMap, uuid) {
			aStar.OpenNodesHeap.Insert(newNode)
		}
	}
}

func worker(id <-chan int, puzzle *Puzzle, aStar *AStarSolver, n *SearchNode, results chan<- *SearchNode) {
	move(ActionsList[<-id], puzzle, aStar, n, results)
}
