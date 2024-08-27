package main

import (
	"fmt"
	"github.com/fatih/color"
)

// Swap two elements in a slice
func Swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Check if a position is within bounds
func InBounds(pos, size int) bool {
	return pos >= 0 && pos < size
}

// Print a formatted puzzle
func PrintFormattedPuzzle(board []int, size int) {
	padding := len(fmt.Sprintf("%d", size*size)) + 1
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if board[x+y*size] == 0 {
				color.New(color.FgRed).Printf("|%*d| ", padding, board[x+y*size])
			} else {
				fmt.Printf("|%*d| ", padding, board[x+y*size])
			}
		}
		fmt.Println()
	}
}
