# Library Management System Documentation

## Overview

The Library Management System is a comprehensive Go-based application designed to efficiently manage books and library members. It allows for the addition and removal of books and members, tracking borrowed books, and generating various reports regarding library inventory and member activity.

## Architecture

The system follows a clean architecture pattern with the following components:

1. **Models**: Define the core data structures used in the application
2. **Services**: Implement the business logic for managing library operations
3. **Controllers**: Handle user interaction and input/output formatting
4. **Utils**: Provide helper functions for input validation and formatting

## Core Components

### Models

#### Book Model
```go
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Status string `json:"status"`
}
```

#### Member Model
```go
type Member struct {
    ID            int    `json:"id"`
    Name          string `json:"name"`
    BorrowedBooks []Book `json:"borrowed_books"`
}
```

#### BorrowedBookRecord Model
```go
type BorrowedBookRecord struct {
    Book     Book   `json:"book"`
    Borrower Member `json:"borrower"`
}
```

### Library Service

The `LibraryService` struct implements the `LibraryManager` interface, which provides the core business logic for the application. It maintains internal storage maps for books, members, and borrowing records:

```go
type LibraryService struct {
    books        map[int]*models.Book
    members      map[int]*models.Member
    borrowedBy   map[int]int
    nextBookID   int
    nextMemberID int
}
```

#### Key Features of the Library Service

1. **Book Management**
   - Add new books with automatic ID assignment
   - Remove books from the collection
   - Check for duplicate books
   - List available books

2. **Member Management**
   - Register new members with automatic ID assignment
   - Remove members from the system
   - Retrieve member information by ID
   - List all registered members

3. **Borrowing System**
   - Track which books are borrowed by which members
   - Handle book borrowing and returning processes
   - List all borrowed books
   - List books borrowed by a specific member

### Library Controller

The controller layer handles user interaction through a command-line interface. It presents menu options, processes user input, and displays formatted output to the user.

#### Key Features of the Library Controller

1. **User Interface**
   - Color-coded menu options for better readability
   - Clear screen functionality for a cleaner interface
   - Formatted output tables for books and member listings

2. **Input Handling**
   - Validation for user inputs
   - Error handling and user-friendly error messages
   - Sanitization of input data

3. **Menu System**
   - Book management menu items (add, remove, borrow, return, list)
   - Member management menu items (add, remove, list)
   - Borrowed books reporting options

## API Reference

### LibraryManager Interface

```go
type LibraryManager interface {
    AddBook(book *models.Book) error
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() ([]*models.Book, error)
    ListBorrowedBooks(memberID int) ([]*models.Book, error)

    AddMember(member *models.Member) (int, error)
    RemoveMember(memberID int) error
    GetMember(memberID int) (*models.Member, error)
    ListMembers() ([]*models.Member, error)
    ListAllBorrowedBooks() ([]*models.BorrowedBookRecord, error)
}
```

#### Book Management Methods

- **AddBook**: Adds a new book to the library collection
  - Parameters: `book *models.Book`
  - Returns: `error` (nil if successful, `ErrDuplicateBook` if the book already exists)

- **RemoveBook**: Removes a book from the library by ID
  - Parameters: `bookID int`
  - Returns: `error` (nil if successful, `ErrBookNotFound` if the book doesn't exist)

- **BorrowBook**: Marks a book as borrowed by a specific member
  - Parameters: `bookID int, memberID int`
  - Returns: `error` (nil if successful, appropriate error if book or member not found)

- **ReturnBook**: Marks a book as returned and available again
  - Parameters: `bookID int, memberID int`
  - Returns: `error` (nil if successful, appropriate error if book or member not found)

- **ListAvailableBooks**: Returns a list of all books with "available" status
  - Returns: `([]*models.Book, error)`

#### Member Management Methods

- **AddMember**: Registers a new member in the system
  - Parameters: `member *models.Member`
  - Returns: `(int, error)` - the member ID and error (nil if successful, `ErrDuplicateMember` if the member exists)

- **RemoveMember**: Removes a member from the system by ID
  - Parameters: `memberID int`
  - Returns: `error` (nil if successful, `ErrMemberNotFound` if the member doesn't exist)

- **GetMember**: Retrieves member information by ID
  - Parameters: `memberID int`
  - Returns: `(*models.Member, error)` - the member object and error

- **ListMembers**: Returns a list of all registered members
  - Returns: `([]*models.Member, error)`

#### Reporting Methods

- **ListBorrowedBooks**: Returns a list of books borrowed by a specific member
  - Parameters: `memberID int`
  - Returns: `([]*models.Book, error)`

- **ListAllBorrowedBooks**: Returns a list of all borrowed books and their borrowers
  - Returns: `([]*models.BorrowedBookRecord, error)`

## Error Handling

The system defines custom errors for common error scenarios:

- `ErrBookNotFound`: Returned when an operation is attempted on a non-existent book
- `ErrDuplicateBook`: Returned when adding a book that already exists
- `ErrMemberNotFound`: Returned when an operation is attempted on a non-existent member
- `ErrDuplicateMember`: Returned when adding a member that already exists

## CLI Interface

The system provides a user-friendly command-line interface with the following options:

1. Book Management:
   - Add Book
   - Remove Book
   - Borrow Book
   - Return Book
   - List Available Books
   - List Borrowed Books

2. Member Management:
   - Add Members
   - Remove Members
   - List Members
   - List Borrowed Books by Member

3. System:
   - Exit

## Usage Example

To initialize and start the Library Management System:

```go
package main

import (
    "intro/task_03/library_management/controllers"
    "intro/task_03/library_management/services"
)

func main() {
    libraryService := services.NewLibraryService()
    libraryController := controllers.NewLibraryController(libraryService)
    libraryController.Start()
}
```

## Conclusion

The Library Management System provides a robust solution for managing library books and members through a simple yet powerful command-line interface. The clean architecture ensures that the codebase is maintainable and extensible, making it easy to add new features or modify existing functionality.
