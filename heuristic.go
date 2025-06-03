package main

import (
	"fmt"
	"runtime"
	"sync"
)

type HeuristicFunction func(board *Puzzle, dt Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
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

type CostFunction func(a, b *SearchNode) int

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
		return aStarCost()
	case uniform:
		return uniformCost()
	}
	return aStarCost()
}

func greedyCost() CostFunction {
	fmt.Println("greedy cost")
	return func(nodeA, nodeB *SearchNode) int {
		return int(*nodeA.HeuristicCost - *nodeB.HeuristicCost)
	}
}

func aStarCost() CostFunction {
	fmt.Println("astar cost")
	return func(nodeA, nodeB *SearchNode) int {
		return int(*nodeA.PathCost+*nodeA.HeuristicCost) - int(*nodeB.PathCost+*nodeB.HeuristicCost)
	}
}

func uniformCost() CostFunction {
	fmt.Println("Uniform cost")
	return func(nodeA, nodeB *SearchNode) int {
		return int(*nodeA.PathCost) - int(*nodeB.PathCost)
	}
}

func ManhattanHeuristic() HeuristicFunction {
	fmt.Println("Manhattan")
	return func(board *Puzzle, final Puzzle) (result int, error error) {
		total := len(board.Tiles)
		workers := runtime.NumCPU()
		if workers > total {
			workers = total
		}
		results := make(chan int, workers)
		var wg sync.WaitGroup
		chunk := (total + workers - 1) / workers
		for w := 0; w < workers; w++ {
			start := w * chunk
			end := start + chunk
			if end > total {
				end = total
			}
			if start >= end {
				continue
			}
			wg.Add(1)
			go func(s, e int) {
				defer wg.Done()
				local := 0
				for i := s; i < e; i++ {
					currentTile := board.Tiles[i]
					finalTile := final.Tiles[i]
					if currentTile != finalTile {
						local += AbsoluteValue(currentTile.X - finalTile.X)
						local += AbsoluteValue(currentTile.Y - finalTile.Y)
					}
				}
				results <- local
			}(start, end)
		}
		wg.Wait()
		close(results)
		for v := range results {
			result += v
		}
		return
	}
}

func VerticalConflict(current, final Tile) (conflicts int) {
	if current.Y == final.Y {
		if current.X != final.X {
			conflicts += AbsoluteValue(current.X - final.X)
		}
	}
	return conflicts * 2
}

func HorizontalConflict(current, final Tile) (conflicts int) {
	if current.X == final.X {
		if current.Y != final.Y {
			conflicts += AbsoluteValue(current.Y - final.Y)
		}
	}
	return conflicts * 2
}

func LinearHeuristic() HeuristicFunction {
	fmt.Println("Manhattan with linear conflicts")
	return func(board *Puzzle, final Puzzle) (result int, error error) {
		total := len(board.Tiles)
		workers := runtime.NumCPU()
		if workers > total {
			workers = total
		}
		results := make(chan int, workers)
		var wg sync.WaitGroup
		chunk := (total + workers - 1) / workers
		for w := 0; w < workers; w++ {
			start := w * chunk
			end := start + chunk
			if end > total {
				end = total
			}
			if start >= end {
				continue
			}
			wg.Add(1)
			go func(s, e int) {
				defer wg.Done()
				local := 0
				for i := s; i < e; i++ {
					currentTile := board.Tiles[i]
					finalTile := final.Tiles[i]
					if currentTile.X != finalTile.X {
						local += AbsoluteValue(currentTile.X - finalTile.X)
					} else {
						local += HorizontalConflict(currentTile, finalTile)
					}
					if currentTile.Y != finalTile.Y {
						local += AbsoluteValue(currentTile.Y - finalTile.Y)
					} else {
						local += VerticalConflict(currentTile, finalTile)
					}
				}
				results <- local
			}(start, end)
		}
		wg.Wait()
		close(results)
		for v := range results {
			result += v
		}
		return
	}
}

func MisplacedHeuristic() HeuristicFunction {
	fmt.Println("Misplaced Tiles")
	return func(board *Puzzle, final Puzzle) (result int, error error) {
		total := len(board.Tiles)
		workers := runtime.NumCPU()
		if workers > total {
			workers = total
		}
		results := make(chan int, workers)
		var wg sync.WaitGroup
		chunk := (total + workers - 1) / workers
		for w := 0; w < workers; w++ {
			start := w * chunk
			end := start + chunk
			if end > total {
				end = total
			}
			if start >= end {
				continue
			}
			wg.Add(1)
			go func(s, e int) {
				defer wg.Done()
				local := 0
				for i := s; i < e; i++ {
					if board.Board[i] != final.Board[i] {
						local++
					}
				}
				results <- local
			}(start, end)
		}
		wg.Wait()
		close(results)
		for v := range results {
			result += v
		}
		return
	}
}
