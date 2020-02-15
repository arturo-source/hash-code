package main

import "github.com/Workiva/go-datastructures/queue"

//Greedy gets the fastes result
func Greedy(maxSlices uint32, pizzasIn []uint32) (pizzasOut []uint32, amount uint32) {
	amount = 0
	pizzasOut = make([]uint32, 0, maxSlices)
	for i, v := range pizzasIn {
		if amount+v > maxSlices {
			return
		}
		amount += v
		pizzasOut = append(pizzasOut, uint32(i))
	}
	return
}

//Recursive gets the best solution
func Recursive(maxSlices uint32, pos uint32, pizzasIn []uint32) (pizzasOut []uint32, amount uint32) {
	if pizzasIn[pos] > maxSlices {
		return
	}
	if pos == uint32(len(pizzasIn))-1 {
		pizzasOut = append(pizzasOut, pos)
		amount = pizzasIn[pos]
		return
	}

	for i := pos; int(i) < len(pizzasIn)-1 && pizzasIn[i] <= maxSlices; i++ {
		tmpPizzas, tmpAmount := Recursive(maxSlices-pizzasIn[i], i+1, pizzasIn)
		tmpAmount += pizzasIn[i]
		if tmpAmount > amount {
			amount = tmpAmount
			pizzasOut = append(tmpPizzas, i)
		}
	}
	return
}

//BranchBound fuck you go compiler
func BranchBound(maxSlices uint32, pizzasIn []uint32) (pizzasOut []uint32, amount uint32) {
	/*
		type TNode struct {
		Pizzas     []uint32
		Pos        uint16
		Deep       uint16
		amount     uint32
		Optimistic uint32
		}
	*/
	pizzasOut, amount = Greedy(maxSlices, pizzasIn)
	initialNode := NewNode(maxSlices, pizzasIn)
	var bestNode *TNode
	pizzaQueue := queue.NewPriorityQueue(0, true)

	pizzaQueue.Put(initialNode)

	for !pizzaQueue.Empty() {
		auxn, _ := pizzaQueue.Get(1)

		nod := auxn[0].(*TNode)

		// if true /*nod.Optimistic > amount*/ {
		if nod.Pessimistic > amount {
			amount = nod.Pessimistic
		}

		if nod.Pessimistic == nod.Amount && nod.Amount >= amount {
			bestNode = nod
		}

		for _, newNod := range nod.Expand(pizzasIn) {
			if newNod.Amount >= amount {
				pizzaQueue.Put(&newNod)
			}
		}
		// }
	}

	if bestNode != nil {
		amount = bestNode.Amount
		pizzasOut = bestNode.Pizzas
	}

	return
}
