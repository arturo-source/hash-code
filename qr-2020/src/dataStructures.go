package main

import (
	"fmt"
	"sort"
)

//LibraryT library structure
type LibraryT struct {
	ID          uint32
	SignTime    uint32
	DailyBooks  uint32
	Books       []uint32
	BooksToSend []uint32
	Score       float32
	MaxBooks    uint32
}

//BookT library structure
type BookT struct {
	// ID    uint32
	Score uint32
	Used  bool
	// Libraries []*LibraryT
}

//Books , all books
var Books []BookT

//Time , total days
var Time uint32

//Libraries all libraries
var Libraries []LibraryT

//LibrariesOrder Indexes of all libraries
var LibrariesOrder []uint32

//CalculateDayValue Calculates the average value of the library per day
func (lib *LibraryT) CalculateDayValue() {
	totalDays := float32(lib.SignTime) + float32(len(lib.Books))/float32(lib.DailyBooks)

	//Calculate value of books
	value := uint32(0)
	for _, bookVal := range lib.Books {
		value += bookVal
	}

	//Calculate average value per day
	lib.Score = float32(value) / float32(totalDays)
	fmt.Println("Score lib", lib.ID, " : ", lib.Score)
}

// func (lib *LibraryT)

//ByScoreBooks to order by score
type ByScoreBooks []uint32

func (a ByScoreBooks) Len() int           { return len(a) }
func (a ByScoreBooks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScoreBooks) Less(i, j int) bool { return Books[a[i]].Score > Books[a[j]].Score }

//SortBooks sort books of a library
func (lib *LibraryT) SortBooks() {
	sort.Sort(ByScoreBooks(lib.Books))
}
