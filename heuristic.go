package main

import (
	"fmt"
)

type HeuristicFunction func(board *Puzzle, dt Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

func FindHeuristic(heuristicType uint) HeuristicFunction {
	fmt.Print("Chosen Heuristic function : ")
	switch heuristicType {
	case Manhattan:
		return ManhattanHeuristic()
	case Linear:
		return LinearHeuristic()
	case Misplaced:
		return MisplacedHeuristic()
	default:
		return ManhattanHeuristic()
	}
}

type CostFunction func(a, b *Node) int

const (
	greedy = iota
	aStar
	uniform
)

func FindCostFunction(costType uint) CostFunction {
	fmt.Print("Chosen Cost function : ")
	switch costType - 1 {
	case greedy:
		return greedyCost()
	case aStar:
		return astarCost()
	case uniform:
		return uniformCost()
	}
	return astarCost()
}

func greedyCost() CostFunction {
	fmt.Println("greedy cost")
	return (func(nodeA, nodeB *Node) int {
		return int(*nodeA.H - *nodeB.H)
	})
}

func astarCost() CostFunction {
	fmt.Println("astar cost")
	return (func(nodeA, nodeB *Node) int {
		return int(*nodeA.G+*nodeA.H) - int(*nodeB.G+*nodeB.H)
	})
}

func uniformCost() CostFunction {
	fmt.Println("Uniform cost")
	return (func(nodeA, nodeB *Node) int {
		return int(*nodeA.G) - int(*nodeB.G)
	})
}

// Add on A the solv function depends on heuristic and fill Solution number
func ManhattanHeuristic() HeuristicFunction {
	fmt.Println("Manhattan")
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (result int, error error) {
		result = 0
		for i := range board.Tiles {
			currentTile := board.Tiles[i]
			finalTile := final.Tiles[i]
			result += Abs(currentTile.X - finalTile.X)
			result += Abs(currentTile.Y - finalTile.Y)
		}
		return
	})
}

func VerticalConflict(current, final Tile) (conflicts int) {
	if current.Y == final.Y {
		if current.X != final.X {
			conflicts += Abs(current.X - final.X)
		}
	}
	return conflicts * 2
}

func HorizontalConflict(current, final Tile) (conflicts int) {
	if current.X == final.X {
		if current.Y != final.Y {
			conflicts += Abs(current.Y - final.Y)
		}
	}
	return conflicts * 2
}

func LinearHeuristic() HeuristicFunction {
	fmt.Println("Manhattan with linear conflicts")
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (result int, error error) {
		result = 0

		for i := range board.Tiles {
			currentTile := board.Tiles[i]
			finalTile := final.Tiles[i]
			if currentTile.X != finalTile.X {
				result += Abs(currentTile.X - finalTile.X)
			} else {
				result += HorizontalConflict(currentTile, finalTile)
			}
			if currentTile.Y != finalTile.Y {
				result += Abs(currentTile.Y - finalTile.Y)
			} else {
				result += VerticalConflict(currentTile, finalTile)
			}
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	fmt.Println("Misplaced Tiles")
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (result int, error error) {
		for i := range board.Tiles {
			if board.Board[i] != final.Board[i] {
				result++
			}
		}
		return
	})
}
