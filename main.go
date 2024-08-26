package main

import (
	"crypto/rand"
	"flag"
	"log"
	"math/big"
)

func main() {
	var puzzleInstance *Puzzle
	commandLineFlags, err := ParseCommandLineFlags()
	if err != nil {
		log.Fatal(err)
	}
	if len(commandLineFlags.Args) == 0 {
		puzzleInstance, err = Generate()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		puzzleInstance, err = LoadPuzzleFromFile(commandLineFlags.Args)
		if err != nil {
			log.Fatal(err)
		}
	}
	Start(puzzleInstance, commandLineFlags.Heuristic-1)
}

type Flags struct {
	Size       int
	Solvable   bool
	Iterations uint
	Cost       uint
	Args       []string
	Heuristic  uint
}

var solvabilityOptions = [2]bool{
	true,
	false,
}

var global Flags

func computeSolv(solvableFlag *bool, solv, unsolv bool) (err error) {
	var randomIndex *big.Int
	if solv {
		*solvableFlag = true
		return nil
	} else if unsolv {
		*solvableFlag = false
		return nil
	} else {
		randomIndex, err = rand.Int(rand.Reader, big.NewInt(int64(len(solvabilityOptions))))
		if err != nil {
			return err
		}
		*solvableFlag = solvabilityOptions[randomIndex.Int64()]
	}
	return nil
}

// Parse func
func ParseCommandLineFlags() (flags Flags, err error) {

	var unsolv bool
	var solv bool
	flag.IntVar(&flags.Size, "size", 3, "Size of the puzzle's side. Must be > 3.")
	flag.BoolVar(&solv, "solvable", false, "Forces generation of a solvable puzzle. Overrides -u.")
	flag.BoolVar(&unsolv, "unsolvable", false, "Forces generation of an unsolvable puzzle.\n(default: random solvable or unsolvable puzzle)")
	flag.UintVar(&flags.Iterations, "iterations", 10000, "Number of iterations.")
	flag.UintVar(&flags.Heuristic, "heu", 1,
		"Forces heuristic, must be between 1 to 3\n\t1 = mahnattan \n\t2 = linear \n\t3 = missplaced \n")
	flag.UintVar(&flags.Cost, "c", 2, "Choose cost, must be between 1 to 3\n\t1 = Greedy Search (Only Heuristic) (faster)\n\t2 = Astar (average)\n\t3 = Uniform search (slower)\n")
	flag.Parse()
	flags.Args = flag.Args()
	if err = computeSolv(&flags.Solvable, solv, unsolv); err != nil {
		return
	}
	if flags.Heuristic < 1 || flags.Heuristic > 3 {
		log.Fatal("Wrong heuristic")
	}
	if flags.Cost < 1 || flags.Cost > 3 {
		log.Fatal("Wrong cost")
	}
	if flags.Size < 3 {
		log.Fatal("Size cant be lower than 3")
	}
	global = flags
	return
}

// Get flags
func GetGlobalFlags() Flags {
	return global
}
