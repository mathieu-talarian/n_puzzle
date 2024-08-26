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

func (p *Puzzle) Move(action int) {
	switch action {
	case ActionTop:
		p.Board.MoveTop(p.Zero.I, p.Size)
	case ActionBot:
		p.Board.MoveBot(p.Zero.I, p.Size)

	case ActionLeft:
		p.Board.MoveLeft(p.Zero.I)

	case ActionRight:
		p.Board.MoveRight(p.Zero.I)
	}

	p.zeroIndex()
	p.TabTiles()
}

func (b *Board) MoveTop(i, size int) {
	tmp := (*b)[i-size]
	(*b)[i-size] = 0
	(*b)[i] = tmp
}

func (b *Board) MoveBot(i, size int) {
	tmp := (*b)[i+size]
	(*b)[i+size] = 0
	(*b)[i] = tmp
}

func (b *Board) MoveLeft(i int) {
	tmp := (*b)[i-1]
	(*b)[i-1] = 0
	(*b)[i] = tmp
}

func (b *Board) MoveRight(i int) {
	tmp := (*b)[i+1]
	(*b)[i+1] = 0
	(*b)[i] = tmp
}

func (p *Puzzle) makeGoals() {
	cur, ix := 1, 1
	x, y, iy := 0, 0, 0
	for {
		(p.Board)[x+y*p.Size] = cur
		if cur == 0 {
			break
		}
		cur++
		if x+ix == p.Size || x+ix < 0 || (ix != 0 && (p.Board)[x+ix+y*p.Size] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == p.Size || y+iy < 0 || (iy != 0 && (p.Board)[x+(y+iy)*p.Size] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cur == p.Size*p.Size {
			cur = 0
		}
	}
}

func initPuzzle(size int) *Puzzle {
	var u = &Puzzle{
		Size:  size,
		Board: make([]int, size*size),
		Tiles: make([]Tile, size*size),
	}
	for i := range u.Board {
		u.Board[i] = -1
	}
	return u
}
func (p *Puzzle) swapEmpty() (err error) {
	if err = p.zeroIndex(); err != nil {
		log.Fatal(err)
	}
	poss := make([]int, 0)
	if (p.Zero.I % p.Size) > 0 {
		poss = append(poss, p.Zero.I-1)
	}
	if (p.Zero.I % p.Size) < (p.Size - 1) {
		poss = append(poss, p.Zero.I+1)
	}
	if (p.Zero.I / p.Size) > 0 {
		poss = append(poss, p.Zero.I-p.Size)
	}
	if (p.Zero.I / p.Size) < (p.Size - 1) {
		poss = append(poss, p.Zero.I+p.Size)
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(poss))))
	swi := poss[n.Int64()]
	p.Board[p.Zero.I], p.Board[swi] = p.Board[swi], 0
	return nil
}

func (p *Puzzle) makePuzzle(solvable bool, iterations uint) (err error) {
	p.makeGoals()
	for i := 0; uint(i) < iterations; i++ {
		if err = p.swapEmpty(); err != nil {
			return
		}
	}
	if !solvable {
		if p.Board[0] == 0 || p.Board[1] == 0 {
			p.Board[len(p.Board)-1], p.Board[len(p.Board)-2] = p.Board[len(p.Board)-2], p.Board[len(p.Board)-1]
		} else {
			p.Board[0], p.Board[1] = p.Board[1], p.Board[0]
		}
	}
	return
}

// ZeroIndex func
func (p *Puzzle) zeroIndex() (err error) {
	for i := range p.Board {
		if p.Board[i] == 0 {
			p.Zero.I = i
			return
		}
	}
	return errors.New("no tile '0'")
}

// Generate function
func Generate() (p *Puzzle, err error) {
	f := Get()
	p = initPuzzle(f.Size)
	if err = p.makePuzzle(f.Solvable, f.Iterations); err != nil {
		return
	}
	if err = p.zeroIndex(); err != nil {
		return
	}
	p.TabTiles()
	return
}

// Tiling returns a tile from a position
func Tiling(size int, pos int) (t Tile) {
	t.X = pos % size
	t.Y = pos / size
	return
}

// TabTiles tiles every puzzle tile
func (p *Puzzle) TabTiles() {
	for i := 0; i < p.Size*p.Size; i++ {
		p.Tiles[p.Board[i]] = Tiling(p.Size, i)
	}
}

// Goal returns a Puzzle and create goal puzzle
func Goal(size int) Puzzle {
	tmp := initPuzzle(size)
	tmp.makeGoals()
	tmp.zeroIndex()
	tmp.TabTiles()
	return *tmp
}

// Inversions returns inversions to calc solvability
func (p *Puzzle) Inversions() (inversions int) {
	for i := range p.Board {
		inversions += inversion(p.Board, i)
	}
	return
}

func inversion(b Board, i int) (inversions int) {
	if b[i] == 0 {
		return 0
	}
	slice := b[i:]
	n := b[i]
	for r := range slice {
		if slice[r] == 0 {
			continue
		}
		if n > slice[r] {
			inversions++
		}
	}
	return inversions
}

// Mod returns
func (p *Puzzle) Mod(i int) int {
	return p.Size % i
}

type TilePosition struct {
	I int
}

// Board is a array of int
type Board []int

// Puzzle struct
type Puzzle struct {
	Zero TilePosition
	Board
	Size int
	Tiles
}

func decompute(str string) *Puzzle {
	t1 := strings.Split(str, "#")
	var board Board
	size, _ := strconv.Atoi(t1[1])
	t2 := strings.Split(t1[0], "|")
	board = make([]int, size*size)
	for i := 0; i < size*size; i++ {
		y, _ := strconv.Atoi(t2[i])
		board[i] = y
	}
	p := &Puzzle{
		Board: board,
		Size:  size,
		Tiles: make([]Tile, size*size),
	}
	p.TabTiles()
	p.zeroIndex()
	return p
}

// CreateUUID builds an uuid from a string
func (p *Puzzle) CreateUUID() TreeString {
	b := p.Board
	tab := make([]string, p.Size*p.Size)
	for k, v := range b {
		tab[k] = strconv.Itoa(v)
	}
	return TreeString(strings.Join(tab, "|"))
}

// Copy deep copy board to board
func (b Board) Copy(i int) Board {
	nb := make([]int, i*i)
	copy(nb, b)
	return nb
}

// Copy deep copy puzzle to puzzle
func (p *Puzzle) Copy() *Puzzle {
	return &Puzzle{
		Zero:  p.Zero,
		Board: p.Board.Copy(p.Size),
		Size:  p.Size,
		Tiles: p.Tiles.Copy(p.Size),
	}
}

// Copy deep copy puzzle tiles to tiles
func (t Tiles) Copy(i int) Tiles {
	np := make([]Tile, i*i)
	copy(np, t)
	return np
}

// PrintPuzzle prints the puzzle on standard input
func (p *Puzzle) PrintPuzzle() {
	u := p.Board
	padding := len(strconv.Itoa(p.Size*p.Size)) + 1
	for y := 0; y < p.Size; y++ {
		for x := 0; x < p.Size; x++ {
			if u[x+y*p.Size] == 0 {
				color.New(color.FgRed).Printf("|%*d| ", padding, u[x+y*p.Size])
			} else {
				fmt.Printf("|%*d| ", padding, u[x+y*p.Size])
			}
		}
		fmt.Printf("\n")
	}
}

// PuzzleFromDatas builds puzzle from an array of int
func PuzzleFromDatas(size int, board []int) (p *Puzzle, err error) {
	p = initPuzzle(size)
	p.Board = board
	p.zeroIndex()
	p.TabTiles()
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
func (t *Tile) TestAction(action int, size int) bool {
	switch action {
	case ActionTop:
		return !(t.Y-1 < 0)
	case ActionBot:
		return t.Y+1 < size
	case ActionLeft:
		return !(t.X-1 < 0)
	case ActionRight:
		return t.X+1 < size
	}
	return false
}

// Bot prints bot
func (t *Tile) Bot() bool {
	fmt.Println("bot")
	return false
}

// Left prints left
func (t *Tile) Left() bool {
	fmt.Println("Left")
	return false
}

// Right prints right
func (t *Tile) Right() bool {
	fmt.Println("right")
	return false
}

// ToTile computes an intex to an x - y  value (tile)
func (t *TilePosition) ToTile(s int) (y *Tile) {
	return &Tile{t.I % s, t.I / s}
}

func (p *Puzzle) compute() string {
	return string(p.CreateUUID()) + "#" + strconv.Itoa(p.Size)
}
