package main

//LibraryT library structure
type LibraryT struct {
	SignTime    uint32
	DailyBooks  uint32
	Books       []uint32
	BooksToSend []uint32
}

//BookT library structure
type BookT struct {
	ID        uint32
	Score     uint32
	Libraries []*LibraryT
}

//Books , all books
var Books []uint32

//Time , total days
var Time uint32

//Libraries all libraries
var Libraries []LibraryT

//LibrariesOrder Indexes of all libraries
var LibrariesOrder []uint32
