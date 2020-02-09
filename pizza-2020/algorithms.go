package main

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
