/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parser.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mmoullec <mmoullec@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/01/15 17:01:22 by mmoullec          #+#    #+#             */
/*   Updated: 2019/01/15 17:01:23 by mmoullec         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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

type Datas struct {
	Size  int
	Board []int
}

func File(av []string) (puzzle *Puzzle, err error) {
	l := list.New()
	if len(av) > 1 {
		return nil, errors.New("Too much arguments")
	}
	file, err := os.Open(av[0])
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newList, err := ScanToList(scanner.Text())
		if err == nil && newList != nil && newList.Len() != 0 {
			l.PushBack(newList)
		}
	}
	d, err := DataFromList(l)
	if err != nil {
		return nil, err
	}
	if len(d.Board) == 0 {
		return nil, fmt.Errorf("Issue with input")
	}
	return PuzzleFromDatas(d.Size, d.Board)
}

func ScanToList(text string) (*list.List, error) {
	tab := strings.Fields(text)
	l := list.New()
	for i := range tab {
		if rune(tab[i][0]) == rune('#') {
			return l, nil
		}
		l.PushBack(tab[i])
	}
	return l, nil
}

func (d *Datas) ListCheckSize(l *list.List) error {
	if l.Len() != 1 {
		return errors.New("Issue with puzzle size")
	} else if s, err := strconv.Atoi(l.Front().Value.(string)); s <= 2 && err == nil {
		return errors.New(fmt.Sprintln("Size too short or negative : ", s))
	} else if err != nil {
		return err
	} else {
		d.Size = s
		d.Board = make([]int, s*s)
		for i := range d.Board {
			d.Board[i] = -1
		}
	}
	return nil
}

func DataFromList(l *list.List) (d *Datas, err error) {
	d = new(Datas)
	i := -1
	cpt := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i == -1 {
			if err = d.ListCheckSize(e.Value.(*list.List)); err != nil {
				return
			}
		} else {
			if l.Len()-1 > d.Size {
				return nil, errors.New("Too much lanes for board")
			}
			if e.Value.(*list.List).Len() != d.Size {
				return nil, errors.New(fmt.Sprintln("Issue with size for lane ", i+1))
			}
			var v int
			for ee := e.Value.(*list.List).Front(); ee != nil; ee = ee.Next() {
				v, err = strconv.Atoi(ee.Value.(string))
				if err != nil {
					return
				}
				if err = CheckNumberIntoBoard(v, d.Size, d.Board); err != nil {
					return nil, err
				}
				d.Board[i] = v
				i++
				cpt++
			}
			i--
		}
		// Si le nombre est superieur ou egal a 0 inferieur a size * size, pas deja utilise ou
		i++
	}
	if cpt < d.Size*d.Size {
		return nil, fmt.Errorf("Issue with puzzle, missing datas")
	}
	return
}

var NumberErr = [2]func(i int) error{
	func(i int) error {
		return errors.New(fmt.Sprintln("Number too low or too high :", i))
	},
	func(i int) error {
		return errors.New(fmt.Sprintln("Number already exists in Board (twice or more):", i))
	},
}

func CheckNumberIntoBoard(n, size int, board []int) error {
	if n < 0 || n >= size*size {
		return NumberErr[0](n)
	}
	cpt := 0
	for _, b := range board {
		if b == n {
			cpt++
		}
	}
	if cpt > 0 {
		return NumberErr[1](n)
	}
	return nil
}
