package main

import (
	"fmt"
	"sync"

	heap "github.com/theodesp/go-heaps"
)

type SearchNode struct {
	MoveAction    *string
	PathCost      *uint
	HeuristicCost *uint
	ParentNode    *SearchNode
	EncodedState  string
}

func NewSearchNode(moveAction *string, pathCost uint, heuristicCost uint, parentNode *SearchNode, puzzleState *Puzzle) *SearchNode {
	return &SearchNode{
		MoveAction:    moveAction,
		PathCost:      &pathCost,
		HeuristicCost: &heuristicCost,
		ParentNode:    parentNode,
		EncodedState:  puzzleState.ComputeEncodedState(),
	}
}

func (node *SearchNode) Compare(otherNode heap.Item) int {
	return costFunction(node, otherNode.(*SearchNode))
}

func (node *SearchNode) IsAlreadyClosed(closedNodesMap map[TreeString]struct{}, nodeUUID TreeString) bool {
	_, isFound := closedNodesMap[nodeUUID]
	return isFound
}

func (node *SearchNode) ExecuteSolver(solver *AStarSolver, nodeUUID TreeString, puzzleState *Puzzle) {
	var wg sync.WaitGroup
	nodeChannel := make(chan *SearchNode, len(ActionsList))
	defer close(nodeChannel)

	for _, action := range ActionsList {
		wg.Add(1)
		go func(action Action) {
			defer wg.Done()
			move(action, puzzleState.Copy(), solver, &node, nodeChannel)
		}(action)
	}

	wg.Wait()

	for range ActionsList {
		add(<-nodeChannel, solver, nodeUUID)
	}
}

func (node *SearchNode) PrintNodeDetails() {
	fmt.Println("Move Action:", *node.MoveAction)
	Decompute(node.EncodedState).PrintPuzzle()
	fmt.Println("Heuristic Cost:", *node.HeuristicCost, "| Path Cost:", *node.PathCost)
	fmt.Println()
}

func (node *SearchNode) PrintSolutionPath() {
	if node != nil {
		node.ParentNode.PrintSolutionPath()
		node.PrintNodeDetails()
	}
}
