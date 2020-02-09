package main

import "fmt"

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
	fmt.Println("im in", pos)
	if amount+pizzasIn[pos] > maxSlices {
		return
	}

	pizzasOut = append(pizzasOut, pos)
	amount = pizzasIn[pos]

	if pos >= uint32(len(pizzasIn))-1 {
		return
	}

	var bestPizzas []uint32
	var bestAmount uint32
	for i := pos; int(i) < len(pizzasIn); i++ {
		tmpPizzas, tmpAmount := Recursive(maxSlices-amount, i+1, pizzasIn)
		tmpAmount += amount
		if tmpAmount > bestAmount && tmpAmount < maxSlices {
			bestAmount = tmpAmount
			bestPizzas = tmpPizzas
		}
	}
	amount = bestAmount
	pizzasOut = append(pizzasOut, bestPizzas...)
	return
}
