package main

import (
	"fmt"
	"log"
)

var costFunction CostFunction

func Start(puzzle *Puzzle, heuristic uint) {
	aStar := NewAStarSolver(puzzle, heuristic)
	costFunction = FindCostFunction(heuristic)
	if !aStar.IsSolvable() {
		log.Fatal("This puzzle is unsolvable")
	}
	if node, err := runN(aStar); err != nil {
		log.Fatal(err)
	} else {
		node.PrintSolutionPath()
		fmt.Println("Number of turns:", aStar.NumberOfTurns)
		fmt.Println("Max state:", aStar.MaxStatesReached)
	}
}

func runN(a *AStarSolver) (q *SearchNode, err error) {
	if err = a.InitializeRootNode(); err != nil {
		return
	}
	for a.OpenNodesHeap.Num() > 0 {
		node := a.OpenNodesHeap.ExtractMinValue()
		state := Decompute(node.(*SearchNode).EncodedState)
		uuid := state.CreateUUID()

		if *node.(*SearchNode).HeuristicCost == 0 {
			return node.(*SearchNode), nil
		}

		(*a).NumberOfTurns++
		node.(*SearchNode).ExecuteSolver(a, uuid, state)
		num := a.OpenNodesHeap.Num()
		if num > a.MaxStatesReached {
			a.MaxStatesReached = num
		}
		a.ClosedNodesMap[uuid] = struct{}{}
	}
	return
}

func move(action Action, state *Puzzle, aStar *AStarSolver, searchNode **SearchNode, results chan<- *SearchNode) {
	tile := state.ZeroPosition.ToTile(state.Size)
	size := state.Size
	if tile.TestAction(action.Value, size) {
		state.Move(action.Value)
		h, err := aStar.HeuristicFunc(state, aStar.GoalPuzzle)
		if err != nil {
			log.Fatal(err)
		}
		results <- NewSearchNode(&action.Name, *(*searchNode).HeuristicCost+1, uint(h), *searchNode, state)
	} else {
		results <- nil
	}
}

func add(newNode *SearchNode, aStar *AStarSolver, uuid TreeString) {
	if newNode != nil {
		if !newNode.IsAlreadyClosed(aStar.ClosedNodesMap, uuid) {
			aStar.OpenNodesHeap.InsertValue(newNode)
		}
	}
}
