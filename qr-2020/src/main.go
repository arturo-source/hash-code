package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello HashCode!")
	if len(os.Args) < 3 {
		panic("There isn't enought arguments")
	}
	filePathIn := os.Args[1]
	filePathOut := os.Args[2]
	fmt.Println(filePathIn, filePathOut)
	libraries, books, time := readFile(filePathIn)
	SortLibraries(libraries, books)
	ReasignBooks(libraries, books, time)
	fmt.Println(books)
	fmt.Println(libraries)
	output(filePathOut, libraries)
}
