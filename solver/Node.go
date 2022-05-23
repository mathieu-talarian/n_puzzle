package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/bst"
	"N_Puzzle/npuzzle"
	"fmt"
	"log"

	heap "github.com/theodesp/go-heaps"
)

type Node struct {
	Action actions.Action
	G      uint
	H      uint
	Parent *Node
	State  npuzzle.Puzzle
}

func (n *Node) Compare(than heap.Item) int {
	return costFunction(n, than.(*Node))
	// return int((n.G + n.H) - (than.(*Node).G + than.(*Node).H))
}

func (n *Node) Tag() interface{} {
	return n.State.CreateUuid()
}

func (n *Node) Key() float64 {
	return float64(n.G + n.H)
}

type INode interface {
	Execute() *Node
}

func NewNode(action actions.Action, g uint, h uint, parent *Node, state npuzzle.Puzzle) *Node {
	return &Node{
		Action: action,
		G:      g,
		H:      h,
		Parent: parent,
		State:  state,
	}
}

func (n *Node) AlreadyClosed(closedList *bst.Node) bool {
	_, ok := closedList.Find(bst.String(n.State.CreateUuid()))
	return ok
}

func move(action actions.Action, state *npuzzle.Puzzle, astar *Astar, n *Node) chan *Node {
	tile := state.Zero.ToTile(state.Size)
	size := state.Size
	ch := make(chan *Node)
	go func() {
		if tile.TestAction(action.Value, size) {
			state.Move(action.Value)
			state.CreateUuid()
			h, err := astar.HeuristicFunction(*state, astar.Goal)
			if err != nil {
				log.Fatal(err)
			}
			ch <- NewNode(action, n.G+1, uint(h), n, *state)
		} else {
			ch <- nil
		}
		close(ch)
	}()
	return ch
}

func add(newNode *Node, a *Astar) {
	if newNode != nil {
		if !Closed_List.AlreadyClosed(newNode) {
			Open_List.Insert(newNode)
		}
	}
}

func (n *Node) Execute(a *Astar) {
	top, bot, left, right := <-move(actions.L[0], n.State.Copy(), a, n), <-move(actions.L[1], n.State.Copy(), a, n), <-move(actions.L[2], n.State.Copy(), a, n), <-move(actions.L[3], n.State.Copy(), a, n)
	add(top, a)
	add(bot, a)
	add(left, a)
	add(right, a)
}

func (n *Node) PrintNode() {
	fmt.Println("Move :", n.Action.Name)
	n.State.PrintPuzzle()
	fmt.Println("Cost:", n.H, "| Depth:", n.G)
	fmt.Println()
}

func (n *Node) PrintResult() {
	if n != nil {
		n.Parent.PrintResult()
		n.PrintNode()
	}
}
