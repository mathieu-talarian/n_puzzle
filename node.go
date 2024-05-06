/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   node.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mmoullec <mmoullec@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/01/15 17:01:14 by mmoullec          #+#    #+#             */
/*   Updated: 2019/01/15 17:01:15 by mmoullec         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"

	heap "github.com/theodesp/go-heaps"
)

/**
Node
*/

type Node struct {
	Action *string
	G      *uint
	H      *uint
	Parent *Node
	State  string
}

func NewNode(action *string, g uint, h uint, parent *Node, state *Puzzle) *Node {
	return &Node{
		Action: action,
		G:      &g,
		H:      &h,
		Parent: parent,
		State:  state.compute(),
	}
}

func (n *Node) Compare(than heap.Item) int {
	return costFunction(n, than.(*Node))
}

func (n *Node) AlreadyClosed(closedList *Bst, uuid BstString) bool {
	ok := closedList.Find(BstString(uuid))
	return ok
}

func (n Node) Execute(a *Astar, uuid BstString, state *Puzzle) {
	id := make(chan int, len(L))
	nodes := make(chan *Node, len(L))
	for range L {
		go worker(id, state.Copy(), a, &n, nodes)
	}
	for _, v := range L {
		id <- v.Value
	}
	close(id)
	for range L {
		add(<-nodes, a, uuid)
	}
	close(nodes)
	id = nil
	nodes = nil
}

func (n *Node) PrintNode() {
	fmt.Println("Move :", *n.Action)
	decompute(n.State).PrintPuzzle()
	fmt.Println("Cost:", *n.H, "| Depth:", *n.G)
	fmt.Println()
}

func (n *Node) PrintResult() {
	if n != nil {
		n.Parent.PrintResult()
		n.PrintNode()
	}
}
