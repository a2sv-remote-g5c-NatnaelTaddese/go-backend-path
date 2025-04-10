package models

import "errors"

var ErrMemberNotFound = errors.New("member not found")
var ErrDuplicateMember = errors.New("duplicate member")

type Member struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BorrowedBooks []Book `json:"borrowed_books"`
}