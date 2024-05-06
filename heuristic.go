package main

import (
	"errors"
	"fmt"
)

type HeuristicFunction func(board *Puzzle, dt Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

func FindHeuristic(h uint) HeuristicFunction {
	fmt.Print("Chosen Heuristic function : ")
	switch h {
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

func FindCostFunction(c uint) CostFunction {
	fmt.Print("Chosen Cost function : ")
	switch c - 1 {
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
	return (func(a, b *Node) int {
		return int(*a.H - *b.H)
	})
}

func astarCost() CostFunction {
	fmt.Println("astar cost")
	return (func(a, b *Node) int {
		return int(*a.G+*a.H) - int(*b.G+*b.H)
	})
}

func uniformCost() CostFunction {
	fmt.Println("Uniform cost")
	return (func(a, b *Node) int {
		return int(*a.G) - int(*b.G)
	})
}

// Add on A the solv function depends on heuristic and fill Solution number
func ManhattanHeuristic() HeuristicFunction {
	fmt.Println("Manhattan")
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			ret += Abs(current.X - final.X)
			ret += Abs(current.Y - final.Y)
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
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			if current.X != final.X {
				ret += Abs(current.X - final.X)
			} else {
				ret += HorizontalConflict(current, final)
			}
			if current.Y != final.Y {
				ret += Abs(current.Y - final.Y)
			} else {
				ret += VerticalConflict(current, final)
			}
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	fmt.Println("Misplaced Tiles")
	return HeuristicFunction(func(board *Puzzle, final Puzzle) (ret int, err error) {
		for i := range board.Tiles {
			if board.Board[i] != final.Board[i] {
				ret++
			}
		}
		return
	})
}
