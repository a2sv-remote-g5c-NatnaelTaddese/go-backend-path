package models

type LibraryManager interface {
	AddBook(book Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error

	ListAvailableBooks() ([]Book, error)
	ListBorrowedBooks(memberID int) ([]Book, error)
}