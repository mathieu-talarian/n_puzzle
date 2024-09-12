package main

import (
	fibonacci "github.com/starwander/GoFibonacciHeap"
)

type AStarSolver struct {
	*Puzzle
	GoalPuzzle       Puzzle
	OpenNodesHeap    fibonacci.FibHeap
	ClosedNodesMap   map[TreeString]struct{}
	NumberOfTurns    uint
	MaxStatesReached uint
	HeuristicFunc    HeuristicFunction
}

func NewAStarSolver(puzzle *Puzzle, heuristicType uint) *AStarSolver {
	return &AStarSolver{
		Puzzle:           puzzle,
		GoalPuzzle:       Goal(puzzle.Size),
		OpenNodesHeap:    *fibonacci.NewFibHeap(),
		ClosedNodesMap:   make(map[TreeString]struct{}),
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
	if err = solver.OpenNodesHeap.InsertValue(NewSearchNode(
		&ActionNone.Name,
		0,
		uint(heuristicValue),
		nil,
		solver.Puzzle)); err != nil {
		return err
	}
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
