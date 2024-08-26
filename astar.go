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

func NewAstar(p *Puzzle, h uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              Goal(p.Size),
		OpenList:          rankparing.New().Init(),
		ClosedList:        nil,
		HeuristicFunction: FindHeuristic(h),
		Turns:             0,
		MaxState:          0,
	}
}

func (a *Astar) RootNode() error {
	h, err := a.HeuristicFunction(a.Puzzle, a.Goal)
	if err != nil {
		return err
	}
	a.OpenList.Insert(NewNode(
		&ActionNone.Name,
		0,
		uint(h),
		nil,
		a.Puzzle))
	return nil
}

func (a *Astar) CheckSolvability() bool {
	a.Puzzle.PrintPuzzle()
	pI := a.Puzzle.Inversions()
	gI := a.Goal.Inversions()
	if a.Puzzle.Mod(2) == 0 {
		pI += a.Puzzle.Zero.I / a.Size
		gI += a.Goal.Zero.I / a.Size
	}
	return pI%2 == gI%2
}
