package controllers

import (
	"bufio"
	"fmt"
	"intro/task_03/library_management/models"
	"intro/task_03/library_management/services"
	"intro/task_03/library_management/utils"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"

	ClearScreen = "\033[H\033[2J"
)

var choice int

type LibraryController struct {
	libraryService *services.LibraryService
	input          bufio.Reader
}

func NewLibraryController(service *services.LibraryService) *LibraryController {
	fmt.Println(Cyan + "\nLibrary Management System Initialized" + Reset)

	return &LibraryController{
		libraryService: service,
		input:          *bufio.NewReader(os.Stdin),
	}
}

func (c *LibraryController) Start() {

	fmt.Println(Green + "Welcome to the Library Management System!" + Reset)

	for {
		fmt.Println(Yellow + "\t1.  Add Book" + Reset)
		fmt.Println(Yellow + "\t2.  Remove Book" + Reset)
		fmt.Println(Yellow + "\t3.  Borrow Book" + Reset)
		fmt.Println(Yellow + "\t4.  Return Book" + Reset)
		fmt.Println(Yellow + "\t5.  List Available Books" + Reset)
		fmt.Println(Yellow + "\t6.  List Borrowed Books\n" + Reset)

		fmt.Println(Purple + "\t7.  Add Members" + Reset)
		fmt.Println(Purple + "\t8.  Remove Members" + Reset)		
		fmt.Println(Purple + "\t9.  List Members" + Reset)
		fmt.Println(Purple + "\t10. List Borrowed Books by Member\n" + Reset)

		fmt.Println(Red + "\t15. Exit" + Reset)

		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		fmt.Print(ClearScreen)
		fmt.Println(Cyan + "Library Management System" + Reset)

		switch choice {

		case 1:
			c.AddBook()

		case 2:
			c.RemoveBook()

		case 3:
			c.BorrowBook()

		case 4:
			c.ReturnBook()

		case 5:
			c.ListAvailableBooks()

		case 6:
			c.ListBorrowedBooks()
		
		case 7:
			c.AddMember()
		
		case 8:
			c.RemoveMember()
		
		case 9:
			c.ListMembers()

		case 10:
			c.ListBorrowedBooksByMember()

		case 15:
			fmt.Println(Green + "Exiting the Library Management System. Goodbye!" + Reset)
			return
		default:
			fmt.Println(Red + "Invalid choice. Please try again." + Reset)
		}
	}
}

func (c *LibraryController) AddBook() {
	fmt.Println(Green + "Adding a new book..." + Reset)

	book := &models.Book{}
	
	// valid title 
	for {
		fmt.Print("Enter book title: ")
		title, err := c.input.ReadString('\n')
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		
		book.Title, err = utils.SanitizeTitle(title)
		if err != nil {
			fmt.Println(Red + "Error in title: " + err.Error() + Reset)
			continue
		}
		break
	}
	
	// valid author
	for {
		fmt.Print("Enter book author: ")
		author, err := c.input.ReadString('\n')
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		
		author = strings.TrimSpace(author)
		if author == "" {
			fmt.Println(Red + "Author cannot be empty. Please try again." + Reset)
			continue
		}
		
		book.Author = author
		break
	}

	book.Status = "available"

	err := c.libraryService.AddBook(book)
	if err != nil {
		fmt.Println(Red + "Error adding book: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "Book added successfully!" + Reset)
	}
}

func (c *LibraryController) RemoveBook() {
	fmt.Println(Green + "Removing a book..." + Reset)

	bookID := 0
	for {
		fmt.Print("Enter book ID to remove: ")
		_, err := fmt.Scan(&bookID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	err := c.libraryService.RemoveBook(bookID)
	if err != nil {
		fmt.Println(Red + "Error removing book: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "Book removed successfully!" + Reset)
	}
}

func (c *LibraryController) BorrowBook() {
	fmt.Println(Green + "Borrowing a book..." + Reset)

	bookID := 0
	memberID := 0

	// valid bookID
	for {
		fmt.Print("Enter book ID to borrow: ")
		_, err := fmt.Scan(&bookID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	// valid memberID
	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	err := c.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(Red + "Error borrowing book: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "Book borrowed successfully!" + Reset)
	}
}

func (c *LibraryController) ReturnBook() {
	fmt.Println(Green + "Returning a book..." + Reset)

	bookID := 0
	memberID := 0

	// valid bookID
	for {
		fmt.Print("Enter book ID to return: ")
		_, err := fmt.Scan(&bookID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	// valid memberID
	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	err := c.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(Red + "Error returning book: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "Book returned successfully!" + Reset)
	}
}

func (c *LibraryController) ListAvailableBooks() {
	fmt.Println(Green + "Listing available books..." + Reset)

	books, err := c.libraryService.ListAvailableBooks()
	if err != nil {
		fmt.Println(Red + "Error listing available books: " + err.Error() + Reset)
		return
	}

	if len(books) == 0 {
		fmt.Println(Yellow + "No available books." + Reset)
		return
	}

	fmt.Println(Cyan + "Available Books:" + Reset)
	fmt.Printf("%-10s %-30s %-30s\n", "ID", "Title", "Author")
	fmt.Println(strings.Repeat("-", 70))
	for _, book := range books {
		fmt.Printf("%-10d %-30s %-30s\n", book.ID, book.Title, book.Author)
	}

		fmt.Println(strings.Repeat("-", 70))
		fmt.Println("\n" + Reset)
}

func (c *LibraryController) ListBorrowedBooks() {
	fmt.Println(Green + "Listing all borrowed books and their borrowers..." + Reset)

	borrowedBooks, err := c.libraryService.ListAllBorrowedBooks()
	if err != nil {
		fmt.Println(Red + "Error listing borrowed books: " + err.Error() + Reset)
		return
	}

	if len(borrowedBooks) == 0 {
		fmt.Println(Yellow + "No borrowed books." + Reset)
		return
	}

	fmt.Println(Cyan + "Borrowed Books:" + Reset)
	fmt.Printf("%-10s %-30s %-30s %-30s\n", "Book ID", "Title", "Author", "Borrower")
	fmt.Println(strings.Repeat("-", 100))
	for _, record := range borrowedBooks {
		fmt.Printf("%-10d %-30s %-30s %-30s\n", record.Book.ID, record.Book.Title, record.Book.Author, record.Borrower.Name)
	}

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\n" + Reset)
}

func (c *LibraryController) AddMember() {
	fmt.Println(Green + "Adding a new member..." + Reset)

	member := &models.Member{}

	// valid name
	for {
		fmt.Print("Enter member name: ")
		name, err := c.input.ReadString('\n')
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println(Red + "Name cannot be empty. Please try again." + Reset)
			continue
		}
		
		member.Name = name
		break
	}

	memberID, err := c.libraryService.AddMember(member)
	if err != nil {
		fmt.Println(Red + "Error adding member: " + err.Error() + Reset)
	} else {
		fmt.Printf(Green + "Member added successfully! Member ID: %d\n" + Reset, memberID)
	}
}

func (c *LibraryController) RemoveMember() {
	fmt.Println(Green + "Removing a member..." + Reset)

	memberID := 0
	for {
		fmt.Print("Enter member ID to remove: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	err := c.libraryService.RemoveMember(memberID)
	if err != nil {
		fmt.Println(Red + "Error removing member: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "Member removed successfully!" + Reset)
	}
}

func (c *LibraryController) ListMembers() {
	fmt.Println(Green + "Listing members..." + Reset)

	members, err := c.libraryService.ListMembers()
	if err != nil {
		fmt.Println(Red + "Error listing members: " + err.Error() + Reset)
		return
	}

	if len(members) == 0 {
		fmt.Println(Yellow + "No members." + Reset)
		return
	}

	fmt.Println(Cyan + "Members:" + Reset)
	fmt.Printf("%-10s %-30s\n", "ID", "Name")
	fmt.Println(strings.Repeat("-", 40))
	for _, member := range members {
		fmt.Printf("%-10d %-30s\n", member.ID, member.Name)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("\n" + Reset)
}

func (c *LibraryController) ListBorrowedBooksByMember() {
	fmt.Println(Green + "Listing borrowed books..." + Reset)

	memberID := 0

	// valid memberID
	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println(Red + "Error reading input: " + err.Error() + Reset)
			continue
		}
		break
	}

	books, err := c.libraryService.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println(Red + "Error listing borrowed books: " + err.Error() + Reset)
		return
	}

	if len(books) == 0 {
		fmt.Println(Yellow + "No borrowed books." + Reset)
		return
	}

	fmt.Println(Cyan + "Borrowed Books:" + Reset)
	fmt.Printf("%-10s %-30s %-30s\n", "ID", "Title", "Author")
	fmt.Println(strings.Repeat("-", 70))
	for _, book := range books {
		fmt.Printf("%-10d %-30s %-30s\n", book.ID, book.Title, book.Author)
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("\n" + Reset)
}