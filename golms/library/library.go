package librarysystem

import (
	"container/list"
	"fmt"
	"library-management-system/book" // Import the book package
	"library-management-system/user" // Import the user package
)

type LibrarySystem struct {
	Users          map[string]string     // Maps user number to user name
	Books          map[string]int        // Maps book name to quantity
	BorrowRegister map[string]*list.List // Maps user number to a list of borrowed book titles
}

// Constructor for LibrarySystem
func NewLibrarySystem() *LibrarySystem {
	return &LibrarySystem{
		Users:          make(map[string]string),
		Books:          make(map[string]int),
		BorrowRegister: make(map[string]*list.List),
	}
}

// Adds a new user
func (l *LibrarySystem) AddUser(u *user.User) {
	l.Users[u.Number] = u.Name
	fmt.Println("User added successfully!")
}

// Adds a new book
func (l *LibrarySystem) AddBook(b *book.Book) {
	l.Books[b.Name] = b.Quantity
	fmt.Println("Book added successfully!")
}

// Views all users
func (l *LibrarySystem) ViewUsers() {
	for number, name := range l.Users {
		fmt.Printf("User: %s, Number: %s\n", name, number)
	}
}

// Views all books
func (l *LibrarySystem) ViewBooks() {
	for name, quantity := range l.Books {
		fmt.Printf("Book: %s, Quantity: %d\n", name, quantity)
	}
}

// Searches for a book by title
func (l *LibrarySystem) SearchBook(name string) (bool, int) {
	quantity, exists := l.Books[name]
	return exists, quantity
}

// Validates a user's existence
func (l *LibrarySystem) IsValidUser(number string) bool {
	_, exists := l.Users[number]
	return exists
}

// Borrow a book
func (l *LibrarySystem) BorrowBook(userNumber, bookTitle string) {
	if !l.IsValidUser(userNumber) {
		fmt.Println("Invalid user!")
		return
	}

	available, quantity := l.SearchBook(bookTitle)
	if !available || quantity == 0 {
		fmt.Println("Book not available!")
		return
	}

	l.Books[bookTitle] -= 1
	if l.BorrowRegister[userNumber] == nil {
		l.BorrowRegister[userNumber] = list.New()
	}
	l.BorrowRegister[userNumber].PushBack(bookTitle)
	fmt.Printf("User %s borrowed book: %s\n", userNumber, bookTitle)
}

// Return a book
func (l *LibrarySystem) ReturnBook(userNumber, bookTitle string) {
	if !l.IsValidUser(userNumber) {
		fmt.Println("Invalid user!")
		return
	}

	borrowedBooks, exists := l.BorrowRegister[userNumber]
	if !exists || borrowedBooks.Len() == 0 {
		fmt.Println("No borrowed books found!")
		return
	}

	for e := borrowedBooks.Front(); e != nil; e = e.Next() {
		if e.Value == bookTitle {
			borrowedBooks.Remove(e)
			l.Books[bookTitle] += 1
			fmt.Printf("Book %s returned by user %s\n", bookTitle, userNumber)
			return
		}
	}
	fmt.Println("This book was not borrowed by the user.")
}

// Displays the borrowed books of a user
func (l *LibrarySystem) DisplayBorrowedBooks(userNumber string) {
	if !l.IsValidUser(userNumber) {
		fmt.Println("Invalid user!")
		return
	}

	borrowedBooks, exists := l.BorrowRegister[userNumber]
	if !exists || borrowedBooks.Len() == 0 {
		fmt.Println("No books borrowed.")
		return
	}

	fmt.Println("Borrowed books:")
	for e := borrowedBooks.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
