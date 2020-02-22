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
	Score uint32
	Used  bool
}

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

//TotalLibsOnTime numero de libros en el tiempo
func (lib *LibraryT) TotalLibsOnTime(time uint32) {
	lib.MaxBooks = uint32((time - lib.SignTime) * lib.DailyBooks)
}

//ByScoreBooks to order by score
type ByScoreBooks struct {
	libBooks   *[]uint32
	totalBooks *[]BookT
}

func (a ByScoreBooks) Len() int { return len((*a.libBooks)) }
func (a ByScoreBooks) Swap(i, j int) {
	(*a.libBooks)[i], (*a.libBooks)[j] = (*a.libBooks)[j], (*a.libBooks)[i]
}
func (a ByScoreBooks) Less(i, j int) bool {
	return (*a.totalBooks)[(*a.libBooks)[i]].Score > (*a.totalBooks)[(*a.libBooks)[j]].Score
}

//SortBooks sort books of a library
func (lib *LibraryT) SortBooks(totalBooks []BookT) {
	sort.Sort(ByScoreBooks{&lib.Books, &totalBooks})
}
