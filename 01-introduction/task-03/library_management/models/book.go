package models

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrDuplicateBook = errors.New("duplicate book")

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
}

type BorrowedBookRecord struct {
	Book     Book   `json:"book"`
	Borrower Member `json:"borrower"`
}