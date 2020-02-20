package main

//LibraryT library structure
type LibraryT struct {
	SignTime    uint32
	DailyBooks  uint32
	Books       []uint32
	BooksToSend []uint32
}

//Books , all books
var Books []uint32

//Time , total days
var Time uint32

//Libraries all libraries
var Libraries []uint32
