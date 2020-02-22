package main

import (
	"sort"
)

//ByScore to order by score
type ByScore []LibraryT

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

//ByScore to order by score
type BySignTime []LibraryT

func (a BySignTime) Len() int           { return len(a) }
func (a BySignTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySignTime) Less(i, j int) bool { return a[i].SignTime < a[j].SignTime }

//SortLibraries Gets de order of the libraries
func SortLibraries(libraries []LibraryT, books []BookT) {
	for i := range libraries {
		libraries[i].CalculateDayValue()
		libraries[i].SortBooks(books)
	}
	sort.Sort(BySignTime(libraries))
	return
}

//ReasignBooks asdf
func ReasignBooks(libraries []LibraryT, books []BookT, time uint32) {
	for i := 0; i < len(libraries); i++ {
		lib := &libraries[i]
		lib.TotalLibsOnTime(time)
		time -= lib.SignTime
		for j := 0; j < int(lib.MaxBooks) && j < len(lib.Books); j++ {
			if !books[lib.Books[j]].Used {
				lib.BooksToSend = append(lib.BooksToSend, lib.Books[j])
				books[lib.Books[j]].Used = true
			}
		}
	}
	return
}
