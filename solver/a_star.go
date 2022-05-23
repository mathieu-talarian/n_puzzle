package solver

import (
	"N_Puzzle/npuzzle"
	"container/list"
	"fmt"
)

type List struct {
	npuzzle.Puzzle
	Next *list.List
}

var costFunction CostFunction

type Astar struct {
	npuzzle.Puzzle
	Goal npuzzle.Puzzle

	MaxState uint
	HeuristicFunction
}

type IAstar interface {
	ManhattanHeuristic() (ret int, err error)
	LinearHeuristic() (ret int, err error)
	MisplacedHeuristic() (ret int, err error)

	Run() (err error)

	RootNode(action int, parent *Node) (err error)

	PrintResult() (err error)

	S()

	Done() bool
}

func (a *Astar) Done() bool {
	return false
}

func NewAstar(p npuzzle.Puzzle, h, c uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              npuzzle.Goal(p.Size),
		HeuristicFunction: FindHeuristic(h),
		MaxState:          0,
	}
}

func (a *Astar) S() {
	fmt.Println("A* =>", a)
}
