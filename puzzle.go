package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func (puzzle *Puzzle) Move(action int) {
	switch action {
	case ActionTop:
		puzzle.Board.MoveTop(puzzle.ZeroPosition.Index, puzzle.Size)
	case ActionBot:
		puzzle.Board.MoveBot(puzzle.ZeroPosition.Index, puzzle.Size)

	case ActionLeft:
		puzzle.Board.MoveLeft(puzzle.ZeroPosition.Index)

	case ActionRight:
		puzzle.Board.MoveRight(puzzle.ZeroPosition.Index)
	}

	puzzle.UpdateZeroIndex()
	puzzle.UpdateTilePositions()
}

func (board *Board) MoveTop(index, size int) {
	temp := (*board)[index-size]
	(*board)[index-size] = 0
	(*board)[index] = temp
}

func (board *Board) MoveBot(index, size int) {
	temp := (*board)[index+size]
	(*board)[index+size] = 0
	(*board)[index] = temp
}

func (board *Board) MoveLeft(index int) {
	temp := (*board)[index-1]
	(*board)[index-1] = 0
	(*board)[index] = temp
}

func (board *Board) MoveRight(index int) {
	temp := (*board)[index+1]
	(*board)[index+1] = 0
	(*board)[index] = temp
}

func (puzzle *Puzzle) CreateGoalState() {
	currentValue, incrementX := 1, 1
	x, y, incrementY := 0, 0, 0
	for {
		(puzzle.Board)[x+y*puzzle.Size] = currentValue
		if currentValue == 0 {
			break
		}
		currentValue++
		if x+incrementX == puzzle.Size || x+incrementX < 0 || (incrementX != 0 && (puzzle.Board)[x+incrementX+y*puzzle.Size] != -1) {
			incrementY = incrementX
			incrementX = 0
		} else if y+incrementY == puzzle.Size || y+incrementY < 0 || (incrementY != 0 && (puzzle.Board)[x+(y+incrementY)*puzzle.Size] != -1) {
			incrementX = -incrementY
			incrementY = 0
		}
		x += incrementX
		y += incrementY
		if currentValue == puzzle.Size*puzzle.Size {
			currentValue = 0
		}
	}
}

func InitializePuzzle(size int) *Puzzle {
	var puzzle = &Puzzle{
		Size:  size,
		Board: make([]int, size*size),
		Tiles: make([]Tile, size*size),
	}
	for i := range puzzle.Board {
		puzzle.Board[i] = -1
	}
	return puzzle
}
func (puzzle *Puzzle) SwapEmptyTile() (err error) {
	if err = puzzle.UpdateZeroIndex(); err != nil {
		log.Fatal(err)
	}
	possibleMoves := make([]int, 0)
	if InBounds(puzzle.ZeroPosition.Index%puzzle.Size-1, puzzle.Size) {
		possibleMoves = append(possibleMoves, puzzle.ZeroPosition.Index-1)
	}
	if InBounds(puzzle.ZeroPosition.Index%puzzle.Size+1, puzzle.Size) {
		possibleMoves = append(possibleMoves, puzzle.ZeroPosition.Index+1)
	}
	if InBounds(puzzle.ZeroPosition.Index/puzzle.Size-1, puzzle.Size) {
		possibleMoves = append(possibleMoves, puzzle.ZeroPosition.Index-puzzle.Size)
	}
	if InBounds(puzzle.ZeroPosition.Index/puzzle.Size+1, puzzle.Size) {
		possibleMoves = append(possibleMoves, puzzle.ZeroPosition.Index+puzzle.Size)
	}
	randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(possibleMoves))))
	swapIndex := possibleMoves[randomIndex.Int64()]
	Swap(puzzle.Board, puzzle.ZeroPosition.Index, swapIndex)
	return nil
}

func (puzzle *Puzzle) GeneratePuzzle(solvable bool, iterations uint) (err error) {
	puzzle.CreateGoalState()
	for i := 0; uint(i) < iterations; i++ {
		if err = puzzle.SwapEmptyTile(); err != nil {
			return
		}
	}
	if !solvable {
		if puzzle.Board[0] == 0 || puzzle.Board[1] == 0 {
			puzzle.Board[len(puzzle.Board)-1], puzzle.Board[len(puzzle.Board)-2] = puzzle.Board[len(puzzle.Board)-2], puzzle.Board[len(puzzle.Board)-1]
		} else {
			puzzle.Board[0], puzzle.Board[1] = puzzle.Board[1], puzzle.Board[0]
		}
	}
	return
}

// UpdateZeroIndex func
func (puzzle *Puzzle) UpdateZeroIndex() (err error) {
	for i := range puzzle.Board {
		if puzzle.Board[i] == 0 {
			puzzle.ZeroPosition.Index = i
			return
		}
	}
	return errors.New("no tile '0'")
}

// Generate function
func Generate() (puzzle *Puzzle, err error) {
	flags := GetGlobalFlags()
	puzzle = InitializePuzzle(flags.Size)
	if err = puzzle.GeneratePuzzle(flags.Solvable, flags.Iterations); err != nil {
		return
	}
	if err = puzzle.UpdateZeroIndex(); err != nil {
		return
	}
	puzzle.UpdateTilePositions()
	return
}

// GetTilePosition returns a tile from a position
func GetTilePosition(size int, position int) (tile Tile) {
	tile.X = position % size
	tile.Y = position / size
	return
}

// UpdateTilePositions updates every puzzle tile position
func (puzzle *Puzzle) UpdateTilePositions() {
	for i := 0; i < puzzle.Size*puzzle.Size; i++ {
		puzzle.Tiles[puzzle.Board[i]] = GetTilePosition(puzzle.Size, i)
	}
}

// Goal returns a Puzzle and create goal puzzle
func Goal(size int) Puzzle {
	tempPuzzle := InitializePuzzle(size)
	tempPuzzle.CreateGoalState()
	tempPuzzle.UpdateZeroIndex()
	tempPuzzle.UpdateTilePositions()
	return *tempPuzzle
}

// Inversions returns inversions to calc solvability
func (puzzle *Puzzle) Inversions() (inversions int) {
	for i := range puzzle.Board {
		inversions += CalculateInversions(puzzle.Board, i)
	}
	return
}

func CalculateInversions(board Board, index int) (inversions int) {
	if board[index] == 0 {
		return 0
	}
	slice := board[index:]
	currentValue := board[index]
	for r := range slice {
		if slice[r] == 0 {
			continue
		}
		if currentValue > slice[r] {
			inversions++
		}
	}
	return inversions
}

// Mod returns
func (puzzle *Puzzle) Mod(i int) int {
	return puzzle.Size % i
}

type TilePosition struct {
	Index int
}

// Board is a array of int
type Board []int

// Puzzle struct
type Puzzle struct {
	ZeroPosition TilePosition
	Board
	Size int
	Tiles
}

func Decompute(encodedString string) *Puzzle {
	parts := strings.Split(encodedString, "#")
	var board Board
	size, _ := strconv.Atoi(parts[1])
	boardValues := strings.Split(parts[0], "|")
	board = make([]int, size*size)
	for i := 0; i < size*size; i++ {
		value, _ := strconv.Atoi(boardValues[i])
		board[i] = value
	}
	puzzle := &Puzzle{
		Board: board,
		Size:  size,
		Tiles: make([]Tile, size*size),
	}
	puzzle.UpdateTilePositions()
	puzzle.UpdateZeroIndex()
	return puzzle
}

// CreateUUID builds an uuid from a string
func (puzzle *Puzzle) CreateUUID() TreeString {
	board := puzzle.Board
	stringValues := make([]string, puzzle.Size*puzzle.Size)
	for index, value := range board {
		stringValues[index] = strconv.Itoa(value)
	}
	return TreeString(strings.Join(stringValues, "|"))
}

// Copy deep copy board to board
func (board *Board) Copy(size int) Board {
	newBoard := make([]int, size*size)
	copy(newBoard, *board)
	return newBoard
}

// Copy deep copy puzzle to puzzle
func (puzzle *Puzzle) Copy() *Puzzle {
	return &Puzzle{
		ZeroPosition: puzzle.ZeroPosition,
		Board:        puzzle.Board.Copy(puzzle.Size),
		Size:         puzzle.Size,
		Tiles:        puzzle.Tiles.Copy(puzzle.Size),
	}
}

// Copy deep copy puzzle tiles to tiles
func (tiles Tiles) Copy(size int) Tiles {
	newTiles := make([]Tile, size*size)
	copy(newTiles, tiles)
	return newTiles
}

// PrintPuzzle prints the puzzle on standard input
func (puzzle *Puzzle) PrintPuzzle() {
	board := puzzle.Board
	padding := len(strconv.Itoa(puzzle.Size*puzzle.Size)) + 1
	for y := 0; y < puzzle.Size; y++ {
		for x := 0; x < puzzle.Size; x++ {
			if board[x+y*puzzle.Size] == 0 {
				color.New(color.FgRed).Printf("|%*d| ", padding, board[x+y*puzzle.Size])
			} else {
				PrintFormattedPuzzle(board, puzzle.Size)
			}
		}
		fmt.Printf("\n")
	}
}

// CreatePuzzleFromDatas builds puzzle from an array of int
func CreatePuzzleFromDatas(size int, board []int) (puzzle *Puzzle, err error) {
	puzzle = InitializePuzzle(size)
	puzzle.Board = board
	puzzle.UpdateZeroIndex()
	puzzle.UpdateTilePositions()
	return
}

// Tile like vector struct X - Y
type Tile struct {
	X int
	Y int
}

// Tiles is an array of tile
type Tiles []Tile

// TestAction switch action kind and returns a bool
func (tile *Tile) TestAction(action int, size int) bool {
	switch action {
	case ActionTop:
		return !(tile.Y-1 < 0)
	case ActionBot:
		return tile.Y+1 < size
	case ActionLeft:
		return !(tile.X-1 < 0)
	case ActionRight:
		return tile.X+1 < size
	}
	return false
}

// Bot prints bot
func (tile *Tile) Bot() bool {
	fmt.Println("bot")
	return false
}

// Left prints left
func (tile *Tile) Left() bool {
	fmt.Println("Left")
	return false
}

// Right prints right
func (tile *Tile) Right() bool {
	fmt.Println("right")
	return false
}

// ToTile computes an intex to an x - y  value (tile)
func (tilePosition *TilePosition) ToTile(size int) (tile *Tile) {
	return &Tile{tilePosition.Index % size, tilePosition.Index / size}
}

func (puzzle *Puzzle) ComputeEncodedState() string {
	return string(puzzle.CreateUUID()) + "#" + strconv.Itoa(puzzle.Size)
}
