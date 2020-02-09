package main

import "fmt"

func main() {
	filePath := "input/e_also_big.in"
	maxSlices, typeSlices, err := read(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Max Slices: ", maxSlices)
	fmt.Println("Content of array: ", typeSlices)
}
