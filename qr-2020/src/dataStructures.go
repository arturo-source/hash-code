package main

//LibraryT library structure
type LibraryT struct {
	ID          uint32
	SignTime    uint32
	DailyBooks  uint32
	Books       []uint32
	BooksToSend []uint32
	Score       uint32
}

//Books , all books
var Books []uint32

//Time , total days
var Time uint32

//Libraries all libraries
var Libraries []LibraryT

//LibrariesOrder Indexes of all libraries
var LibrariesOrder []uint32
