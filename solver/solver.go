package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"fmt"
	"log"
)

// Start function init astar
func Start(p npuzzle.Puzzle, h uint, c uint) {
	a := NewAstar(p, h, c)
	costFunction = FindCostFunction(c)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	fmt.Println("Searching solution...")
	if n, err := a.Run(p.Size); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
		fmt.Println("Number of turns:", Turns.val)
		fmt.Println("Max state:", a.MaxState)
	}
}

const (
	//No action
	No = iota
)

// Run function Runs the astar algorithm
func (a *Astar) Run(size int /*FCost SortList */) (q *Node, err error) {
	//if size < 4 {
	//	return run(a)
	//}
	return runN(a)
}

func runN(a *Astar /* , FCost */) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}

	for true {
		for i := 0; i < int(NUM_JOBS); i++ {
			Turns.Inc()
			Jobs <- a
		}
		for i := 0; i < int(NUM_JOBS); i++ {
			res := <-Results
			if res != nil {
				return res, nil
			}
		}
	}

	return
}

/*
RootNode func
*/
func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	if err != nil {
		return
	}
	Open_List.Insert(
		NewNode(
			actions.None,
			0,
			uint(h),
			nil,
			a.Puzzle))
	return
}
