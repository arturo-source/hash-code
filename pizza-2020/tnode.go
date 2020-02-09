package main

type TNode struct {
	Pizzas     []uint32
	Pos        uint16
	Deep       uint16
	amount     uint32
	Optimistic uint32
}

func (tnode *TNode) Expand(pizzasIn []uint32) (sons []TNode) {
	return
}

func (tnode *TNode) Feasible() (sons []TNode) {
	return
}
