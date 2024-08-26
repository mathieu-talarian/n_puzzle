package main

import (
	rankparing "github.com/theodesp/go-heaps/rank_pairing"
)

type Astar struct {
	*Puzzle
	Goal       Puzzle
	OpenList   *rankparing.RPHeap
	ClosedList *Bst
	Turns      uint
	MaxState   uint
	HeuristicFunction
}

func NewAstar(puzzle *Puzzle, heuristic uint) *Astar {
	return &Astar{
		Puzzle:            puzzle,
		Goal:              Goal(puzzle.Size),
		OpenList:          rankparing.New().Init(),
		ClosedList:        nil,
		HeuristicFunction: FindHeuristic(heuristic),
		Turns:             0,
		MaxState:          0,
	}
}

func (astar *Astar) RootNode() error {
	heuristic, err := astar.HeuristicFunction(astar.Puzzle, astar.Goal)
	if err != nil {
		return err
	}
	astar.OpenList.Insert(NewNode(
		&ActionNone.Name,
		0,
		uint(h),
		nil,
		astar.Puzzle))
	return nil
}

func (astar *Astar) CheckSolvability() bool {
	astar.Puzzle.PrintPuzzle()
	puzzleInversions := astar.Puzzle.Inversions()
	goalInversions := astar.Goal.Inversions()
	if astar.Puzzle.Mod(2) == 0 {
		puzzleInversions += astar.Puzzle.Zero.I / astar.Size
		goalInversions += astar.Goal.Zero.I / astar.Size
	}
	return puzzleInversions%2 == goalInversions%2
}
