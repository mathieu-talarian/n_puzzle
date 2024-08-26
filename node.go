package main

import (
	"fmt"

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

func (node *SearchNode) IsAlreadyClosed(closedNodesTree *BinarySearchTree, nodeUUID TreeString) bool {
	isFound := closedNodesTree.Find(TreeString(nodeUUID))
	return isFound
}

func (node SearchNode) ExecuteSolver(solver *AStarSolver, nodeUUID TreeString, puzzleState *Puzzle) {
	actionChannel := make(chan int, len(ActionsList))
	nodeChannel := make(chan *SearchNode, len(ActionsList))
	defer close(actionChannel)
	defer close(nodeChannel)
	for range ActionsList {
		go worker(actionChannel, puzzleState.Copy(), solver, &node, nodeChannel)
	}
	for _, action := range ActionsList {
		actionChannel <- action.Value
	}
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
