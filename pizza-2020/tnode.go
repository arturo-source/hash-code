package main

import (
	"github.com/Workiva/go-datastructures/queue"
)

//TNode structure
type TNode struct {
	Pizzas      []uint16
	Pos         uint16
	Deep        uint16
	Amount      uint32
	Pessimistic uint32
	Optimistic  uint32
	PizzasPess  []uint16
}

func (tnode *TNode) Compare(other queue.Item) int {
	otherN := other.(*TNode)
	if tnode.Pessimistic < otherN.Pessimistic {
		return 1
	}
	return -1
}

func (tnode *TNode) setPessimistic(maxSlices uint32, pizzasIn []uint32) {
	tnode.Pessimistic = tnode.Amount
	tnode.PizzasPess = tnode.Pizzas
	for i := tnode.Pos + 1; i < uint16(len(pizzasIn)); i++ {
		slices := pizzasIn[i]
		if tnode.Pessimistic+slices > maxSlices {
			return
		}
		tnode.Pessimistic += slices
		tnode.PizzasPess = append(tnode.PizzasPess, i)
	}
	return
}

//setOptimistic sets the optimistic bound fuck you go compiler
func (tnode *TNode) setOptimistic(maxSlices uint32, pizzasIn []uint32) {
	if tnode.Pos == uint16(len(pizzasIn)-1) {
		tnode.Optimistic = tnode.Amount
		return
	}
	min := pizzasIn[tnode.Pos+1]
	max := pizzasIn[len(pizzasIn)-1]

	actualValue := maxSlices - tnode.Amount

	if actualValue <= max && actualValue >= min {
		tnode.Optimistic = actualValue + tnode.Amount
		return
	}

	var bestValue uint32 = 0

	for i := min; i <= max; i++ {
		divResult := actualValue / i
		divModule := actualValue % i

		if divModule == 0 || divModule >= min {
			tnode.Optimistic = actualValue + tnode.Amount
			return
		}

		if divResult > bestValue {
			bestValue = divResult
		}
	}

	tnode.Optimistic = bestValue + tnode.Amount

	return
}

func NewNode(maxSlices uint32, pizzasIn []uint32) (tnode *TNode) {
	if len(pizzasIn) > 0 {
		tnode = &TNode{
			[]uint16{0},
			0,
			0,
			0,
			0,
			0,
			[]uint16{}}
	}
	tnode.setPessimistic(maxSlices, pizzasIn)
	tnode.setOptimistic(maxSlices, pizzasIn)
	return
}

func (tnode *TNode) Expand(maxSlices uint32, pizzasIn []uint32) (sons []TNode) {
	// fmt.Println("Father ", *tnode)
	for i := tnode.Pos + 1; i < uint16(len(pizzasIn)); i++ {
		amount := tnode.Amount + pizzasIn[i]
		if amount <= maxSlices {
			auxPizzas := make([]uint16, len(tnode.Pizzas), len(tnode.Pizzas)+1)
			copy(auxPizzas, tnode.Pizzas)
			auxNode := TNode{append(auxPizzas, i), i, tnode.Deep + 1, amount, 0, 0, []uint16{}}
			auxNode.setPessimistic(maxSlices, pizzasIn)
			auxNode.setOptimistic(maxSlices, pizzasIn)
			// fmt.Println("Son ", auxNode)
			sons = append(sons, auxNode)
		}
	}
	return
}
