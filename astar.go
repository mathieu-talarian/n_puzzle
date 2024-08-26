package main

import (
	rankparing "github.com/theodesp/go-heaps/rank_pairing"
)

type AStarSolver struct {
	*Puzzle
	GoalPuzzle       Puzzle
	OpenNodesHeap    *rankparing.RPHeap
	ClosedNodesTree  *BinarySearchTree
	NumberOfTurns    uint
	MaxStatesReached uint
	HeuristicFunc    HeuristicFunction
}

func NewAStarSolver(puzzle *Puzzle, heuristicType uint) *AStarSolver {
	return &AStarSolver{
		Puzzle:           puzzle,
		GoalPuzzle:       Goal(puzzle.Size),
		OpenNodesHeap:    rankparing.New().Init(),
		ClosedNodesTree:  nil,
		HeuristicFunc:    FindHeuristic(heuristicType),
		NumberOfTurns:    0,
		MaxStatesReached: 0,
	}
}

func (solver *AStarSolver) InitializeRootNode() error {
	heuristicValue, err := solver.HeuristicFunc(solver.Puzzle, solver.GoalPuzzle)
	if err != nil {
		return err
	}
	solver.OpenNodesHeap.Insert(NewSearchNode(
		&ActionNone.Name,
		0,
		uint(heuristicValue),
		nil,
		solver.Puzzle))
	return nil
}

func (solver *AStarSolver) IsSolvable() bool {
	solver.Puzzle.PrintPuzzle()
	puzzleInversions := solver.Puzzle.Inversions()
	goalInversions := solver.GoalPuzzle.Inversions()
	if solver.Puzzle.Mod(2) == 0 {
		puzzleInversions += solver.Puzzle.ZeroPosition.Index / solver.Size
		goalInversions += solver.GoalPuzzle.ZeroPosition.Index / solver.Size
	}
	return puzzleInversions%2 == goalInversions%2
}
