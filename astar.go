package main

import (
	rank_paring "github.com/theodesp/go-heaps/rank_pairing"
)

type Astar struct {
	*Puzzle
	Goal       Puzzle
	OpenList   *rank_paring.RPHeap
	ClosedList *Bst
	Turns      uint
	MaxState   uint
	HeuristicFunction
}

func NewAstar(p *Puzzle, h, c uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              Goal(p.Size),
		OpenList:          rank_paring.New().Init(),
		ClosedList:        nil,
		HeuristicFunction: FindHeuristic(h),
		Turns:             0,
		MaxState:          0,
	}
}

func (a *Astar) RootNode(action int) error {
	h, err := a.HeuristicFunction(a.Puzzle, a.Goal)
	if err != nil {
		return err
	}
	a.OpenList.Insert(NewNode(
		&None.Name,
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
