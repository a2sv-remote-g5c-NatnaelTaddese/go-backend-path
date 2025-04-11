package services

import "intro/task_03/library_management/models"

type LibraryService struct {
	books        map[int]*models.Book
	members      map[int]*models.Member
	borrowedBy   map[int]int
	nextBookID   int
	nextMemberID int
}

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

func NewLibraryService() LibraryManager {
	return &LibraryService{
		books:        make(map[int]*models.Book),
		members:      make(map[int]*models.Member),
		borrowedBy:   make(map[int]int),
		nextBookID:   1,
		nextMemberID: 1,
	}
}

func (s *LibraryService) AddBook(book *models.Book) error {
	for _, existingBook := range s.books {
		if existingBook.Title == book.Title && existingBook.Author == book.Author {
			return models.ErrDuplicateBook
		}
	}
	book.ID = s.nextBookID
	s.books[s.nextBookID] = book
	s.nextBookID++
	return nil
}

func (s *LibraryService) RemoveBook(bookID int) error {
	if _, exists := s.books[bookID]; !exists {
		return models.ErrBookNotFound
	}
	delete(s.books, bookID)
	return nil
}

func (s *LibraryService) BorrowBook(bookID int, memberID int) error {
	if _, exists := s.books[bookID]; !exists {
		return models.ErrBookNotFound
	}
	if _, exists := s.members[memberID]; !exists {
		return models.ErrMemberNotFound
	}
	s.borrowedBy[bookID] = memberID
	book := s.books[bookID]
	book.Status = "borrowed"
	s.books[bookID] = book
	return nil
}

func (s *LibraryService) ReturnBook(bookID int, memberID int) error {
	if _, exists := s.books[bookID]; !exists {
		return models.ErrBookNotFound
	}
	if _, exists := s.members[memberID]; !exists {
		return models.ErrMemberNotFound
	}
	if _, exists := s.borrowedBy[bookID]; !exists {
		return models.ErrBookNotFound
	}
	delete(s.borrowedBy, bookID)
	book := s.books[bookID]
	book.Status = "available"
	s.books[bookID] = book
	return nil
}

func (s *LibraryService) ListAvailableBooks() ([]*models.Book, error) {
	var availableBooks []*models.Book
	for _, book := range s.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks, nil
}

func (s *LibraryService) ListBorrowedBooks(memberID int) ([]*models.Book, error) {
	var borrowedBooks []*models.Book
	for bookID, borrowerID := range s.borrowedBy {
		if borrowerID == memberID {
			borrowedBooks = append(borrowedBooks, s.books[bookID])
		}
	}
	return borrowedBooks, nil
}

func (s *LibraryService) AddMember(member *models.Member) (int, error) {
	for _, existingMember := range s.members {
		if existingMember.Name == member.Name {
			return existingMember.ID, models.ErrDuplicateMember
		}
	}
	member.ID = s.nextMemberID
	s.members[s.nextMemberID] = member
	s.nextMemberID++
	return member.ID, nil
}

func (s *LibraryService) RemoveMember(memberID int) error {
	if _, exists := s.members[memberID]; !exists {
		return models.ErrMemberNotFound
	}
	delete(s.members, memberID)
	return nil
}

func (s *LibraryService) GetMember(memberID int) (*models.Member, error) {
	if member, exists := s.members[memberID]; exists {
		return member, nil
	}
	return nil, models.ErrMemberNotFound
}

func (s *LibraryService) ListMembers() ([]*models.Member, error) {
	var membersList []*models.Member
	for _, member := range s.members {
		membersList = append(membersList, member)
	}
	return membersList, nil
}

func (s *LibraryService) ListAllBorrowedBooks() ([]*models.BorrowedBookRecord, error) {
	var borrowedBooks []*models.BorrowedBookRecord
	for bookID, memberID := range s.borrowedBy {
		book, exists := s.books[bookID]
		if !exists {
			continue
		}

		member, exists := s.members[memberID]
		if !exists {
			continue
		}

		borrowedBooks = append(borrowedBooks, &models.BorrowedBookRecord{
			Book:     *book,
			Borrower: *member,
		})
	}
	return borrowedBooks, nil
}
