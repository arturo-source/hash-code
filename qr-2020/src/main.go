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
	libraries, books, time := readFile(filePathIn)
	SortLibraries(libraries)
	ReasignBooks(libraries, books, time)
	fmt.Println(books)
	fmt.Println(libraries)
	output(filePathOut, libraries)
}

func readFile(fName string) (libraries []LibraryT, books []BookT, time uint32) {
	dat, err := ioutil.ReadFile(fName)
	check(err)

	//Read total Books libraries and Days to scan
	fileLines := strings.Split(string(dat), "\n")

	aux := strings.Split(fileLines[0], " ")
	totalBooks, _ := strconv.Atoi(aux[0])
	totalLibs, _ := strconv.Atoi(aux[1])
	auxTime, _ := strconv.Atoi(aux[2])
	time = uint32(auxTime)

	//Read books
	books = make([]BookT, totalBooks)
	auxBooks := strings.Split(fileLines[1], " ")
	for i := 0; i < totalBooks; i++ {
		auxValue, _ := strconv.Atoi(auxBooks[i])
		var book BookT
		book.Score = uint32(auxValue)
		book.Used = false
		books[i] = book
	}

	//Read libraries
	libraries = make([]LibraryT, totalLibs)
	length := totalLibs * 2
	for i := 2; i <= length; i += 2 {
		//Read library initial data
		auxLibs := strings.Split(fileLines[i], " ")
		var lib LibraryT

		totalBooksLib, _ := strconv.Atoi(auxLibs[0])
		singupTime, _ := strconv.Atoi(auxLibs[1])
		shipLimit, _ := strconv.Atoi(auxLibs[2])

		lib.Books = make([]uint32, 0, totalBooksLib)
		lib.SignTime = uint32(singupTime)
		lib.DailyBooks = uint32(shipLimit)

		//Read library books information
		for _, val := range strings.Split(fileLines[i+1], " ") {
			bookIndex, _ := strconv.Atoi(val)
			lib.Books = append(lib.Books, uint32(bookIndex))
		}
		calcPos := uint32(i/2 - 1)
		lib.ID = calcPos
		libraries[calcPos] = lib
	}

	fmt.Println("Days: ", time)
	fmt.Println("Books: ", totalBooks)
	fmt.Println(books)
	fmt.Println("Libraries: ", totalLibs)
	fmt.Println(libraries)
	return
}
