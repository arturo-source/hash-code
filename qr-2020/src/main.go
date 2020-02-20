package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	readFile(filePathIn)
}

func readFile(fName string) {
	dat, err := ioutil.ReadFile(fName)
	check(err)

	//Read total Books Libraries and Days to scan
	fileLines := strings.Split(string(dat), "\n")

	aux := strings.Split(fileLines[0], " ")
	totalBooks, _ := strconv.Atoi(aux[0])
	totalLibs, _ := strconv.Atoi(aux[1])
	auxTime, _ := strconv.Atoi(aux[2])
	Time = uint32(auxTime)

	//Read books
	Books = make([]uint32, totalBooks)
	auxBooks := strings.Split(fileLines[1], " ")
	for i := 0; i < totalBooks; i++ {
		auxValue, _ := strconv.Atoi(auxBooks[i])
		Books[i] = uint32(auxValue)
	}

	//Read Libraries
	Libraries = make([]LibraryT, totalLibs)
	length := totalLibs * 2
	for i := 2; i <= length; i += 2 {
		//Read library initial data
		auxLibs := strings.Split(fileLines[i], " ")
		var lib LibraryT

		totalBooksLib, _ := strconv.Atoi(auxLibs[0])
		singupTime, _ := strconv.Atoi(auxLibs[1])
		shipLimit, _ := strconv.Atoi(auxLibs[2])

		lib.Books = make([]uint32, totalBooksLib)
		lib.SignTime = uint32(singupTime)
		lib.DailyBooks = uint32(shipLimit)

		//Read library books information
		for j, val := range strings.Split(fileLines[i+1], " ") {
			auxValue, _ := strconv.Atoi(val)
			lib.Books[j] = uint32(auxValue)
		}
		calcPos := uint32(i/2 - 1)
		lib.ID = calcPos
		lib.CalculateDayValue()
		Libraries[calcPos] = lib
	}

	fmt.Println("Days: ", Time)
	fmt.Println("Books: ", totalBooks)
	fmt.Println(Books)
	fmt.Println("Libraries: ", totalLibs)
	fmt.Println(Libraries)
}
