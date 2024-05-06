package main

import (
	"fmt"

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

func (a *Astar) Done() bool {
	return false
}

func (a *Astar) S() {
	fmt.Println("A* =>", a)
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

/*
RootNode func
*/
func (a *Astar) RootNode(action int) (err error) {
	var h int
	h, err = a.HeuristicFunction(a.Puzzle, a.Goal)
	if err != nil {
		return
	}
	a.OpenList.Insert(NewNode(
		&None.Name,
		0,
		uint(h),
		nil,
		a.Puzzle))
	return
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
