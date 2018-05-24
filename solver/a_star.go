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

type Astar struct {
	npuzzle.Puzzle
	Goal       npuzzle.Puzzle
	OpenList   *list.List
	ClosedList *list.List
	Turns      uint
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
}

func NewAstar(p npuzzle.Puzzle, h uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              npuzzle.Goal(p.Size),
		OpenList:          list.New(),
		ClosedList:        list.New(),
		HeuristicFunction: FindHeuristic(h),
		Turns:             0,
	}
}

func (a *Astar) S() {
	fmt.Println("A* =>", a)
}