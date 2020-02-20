package main

import "sort"

//ByScore to order by score
type ByScore []LibraryT

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

//SortLibraries Gets de order of the libraries
func SortLibraries(libraries []LibraryT) {
	sort.Sort(ByScore(libraries))
	return
}

//ReasignBooks asdf
func ReasignBooks(libraries []LibraryT) {
	for i := 0; i < len(libraries); i++ {
		lib := &libraries[i]
		for j := 0; j < int(lib.MaxBooks) && j < len(lib.Books); j++ {
			if !Books[lib.Books[j]].Used {
				lib.BooksToSend = append(lib.BooksToSend, lib.Books[j])
				Books[lib.Books[j]].Used = true
			}
		}
	}
	return
}
