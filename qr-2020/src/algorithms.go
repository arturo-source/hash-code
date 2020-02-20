package main

import "sort"

//SortBooks sort books of a library
func (lib *LibraryT) SortBooks() {

}

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
