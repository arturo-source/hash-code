package main

import (
	"github.com/Workiva/go-datastructures/queue"
)

type TNode struct {
	Pizzas      []uint16
	Pos         uint16
	Deep        uint16
	Amount      uint32
	Pessimistic uint32
}

func (tnode *TNode) Compare(queue.Item) int {
	return 0
}

func (tnode *TNode) setPessimistic(maxSlices uint32, pizzasIn []uint32) {
	tnode.Pessimistic = tnode.Amount
	for i := tnode.Pos; i < uint16(len(pizzasIn)); i++ {
		slices := pizzasIn[i]
		if tnode.Pessimistic+slices > maxSlices {
			return
		}
		tnode.Pessimistic += slices
	}
	return
}

func NewNode(maxSlices uint32, pizzasIn []uint32) (tnode *TNode) {
	if len(pizzasIn) > 0 {
		tnode = &TNode{
			[]uint16{0},
			0,
			0,
			0,
			0}
	}
	tnode.setPessimistic(maxSlices, pizzasIn)
	return
}

func (tnode *TNode) Expand(maxSlices uint32, pizzasIn []uint32) (sons []TNode) {
	for i := tnode.Pos + 1; i < uint16(len(pizzasIn)); i++ {
		amount := tnode.Amount + pizzasIn[i]
		if amount <= maxSlices {
			auxNode := TNode{append(tnode.Pizzas, i), i, tnode.Deep + 1, amount, 0}
			auxNode.setPessimistic(maxSlices, pizzasIn)
			sons = append(sons, auxNode)
		}
	}
	return
}
