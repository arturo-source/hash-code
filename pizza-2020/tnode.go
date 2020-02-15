package main

import (
	"github.com/Workiva/go-datastructures/queue"
)

type TNode struct {
	Pizzas     []uint32
	Pos        uint16
	Deep       uint16
	Amount     uint32
	Optimistic uint32
}

func (tnode *TNode) Compare(queue.Item) int {
	return 0
}

func (tnode *TNode) setOptimistic(maxSlices uint32, pizzasIn []uint32) {
	tnode.Optimistic = tnode.Amount
	for i := tnode.Pos; i < uint16(len(pizzasIn)); i++ {
		slices := pizzasIn[i]
		if tnode.Optimistic+slices > maxSlices {
			return
		}
		tnode.Optimistic += slices
	}
	return
}

func NewNode(maxSlices uint32, pizzasIn []uint32) (tnode *TNode) {
	if len(pizzasIn) > 0 {
		tnode = &TNode{
			[]uint32{0},
			0,
			0,
			0,
			0}
	}
	tnode.setOptimistic(maxSlices, pizzasIn)
	return
}

func (tnode *TNode) Expand(pizzasIn []uint32) (sons []TNode) {
	for i := tnode.Pos + 1; i < uint16(len(pizzasIn)); i++ {

	}
	return
}

func (tnode *TNode) Feasible() (sons []TNode) {
	return
}
