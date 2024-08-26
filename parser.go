package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PuzzleConfiguration struct {
	PuzzleSize  int
	PuzzleBoard []int
}

func LoadPuzzleFromFile(filePaths []string) (puzzle *Puzzle, err error) {
	linesList := list.New()
	if len(filePaths) > 1 {
		return nil, errors.New("too much arguments")
	}
	file, err := os.Open(filePaths[0])
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesList.Init()
		parsedList, err := ParseLineToList(scanner.Text(), linesList)
		if err == nil && parsedList != nil && parsedList.Len() != 0 {
			linesList.PushBack(parsedList)
		}
	}
	puzzleData, err := ExtractDataFromList(linesList)
	if err != nil {
		return nil, err
	}
	if len(puzzleData.PuzzleBoard) == 0 {
		return nil, fmt.Errorf("issue with input")
	}
	return CreatePuzzleFromDatas(puzzleData.PuzzleSize, puzzleData.PuzzleBoard)
}

func ParseLineToList(line string, list *list.List) (*list.List, error) {
	fields := strings.Fields(line)
	for i := range fields {
		if strings.HasPrefix(fields[i], "#") {
			return list, nil
		}
		list.PushBack(fields[i])
	}
	return list, nil
}

func (puzzleData *PuzzleConfiguration) ValidateListSize(list *list.List) error {
	if list.Len() != 1 {
		return errors.New("issue with puzzle size")
	} else if size, err := strconv.Atoi(list.Front().Value.(string)); size <= 2 && err == nil {
		return errors.New(fmt.Sprintln("Size too short or negative : ", size))
	} else if err != nil {
		return err
	} else {
		puzzleData.PuzzleSize = size
		puzzleData.PuzzleBoard = make([]int, size*size)
		for i := range puzzleData.PuzzleBoard {
			puzzleData.PuzzleBoard[i] = -1
		}
	}
	return nil
}

func ExtractDataFromList(list *list.List) (puzzleData *PuzzleConfiguration, err error) {
	puzzleData = new(PuzzleConfiguration)
	index := -1
	count := 0
	for element := list.Front(); element != nil; element = element.Next() {
		if index == -1 {
			if err = puzzleData.ValidateListSize(element.Value.(*list.List)); err != nil {
				return
			}
		} else {
			if list.Len()-1 > puzzleData.PuzzleSize {
				return nil, errors.New("too much lanes for board")
			}
			if element.Value.(*list.List).Len() != puzzleData.PuzzleSize {
				return nil, errors.New(fmt.Sprintln("Issue with size for lane ", index+1))
			}
			var value int
			for subElement := element.Value.(*list.List).Front(); subElement != nil; subElement = subElement.Next() {
				value, err = strconv.Atoi(subElement.Value.(string))
				if err != nil {
					return
				}
				if err = ValidateNumberInBoard(value, puzzleData.PuzzleSize, puzzleData.PuzzleBoard); err != nil {
					return nil, err
				}
				puzzleData.PuzzleBoard[index] = value
				index++
				count++
			}
			index--
		}
		index++
	}
	if count < puzzleData.PuzzleSize*puzzleData.PuzzleSize {
		return nil, fmt.Errorf("issue with puzzle, missing data")
	}
	return
}

func ValidateNumberInBoard(number, size int, board []int) error {
	if number < 0 || number >= size*size {
		return errors.New(fmt.Sprintln("Number too low or too high :", number))

	}
	count := 0
	for _, boardNumber := range board {
		if boardNumber == number {
			count++
		}
	}
	if count > 0 {
		return errors.New(fmt.Sprintln("Number already exists in Board (twice or more):", number))
	}
	return nil
}
