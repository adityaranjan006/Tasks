package main

import (
	"fmt"
	"library-management-system/book"
	librarysystem "library-management-system/library"
	"library-management-system/user"
	"time"
)

func main() {
	lms := librarysystem.NewLibrarySystem()

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Users")
		fmt.Println("2. Add Books")
		fmt.Println("3. View Books")
		fmt.Println("4. View Users")
		fmt.Println("5. Borrow Book")
		fmt.Println("6. Return Book")
		fmt.Println("7. View User Bucket")
		fmt.Println("8. Exit Application")
		fmt.Print("Enter an option: ")

		var option int
		_, err := fmt.Scanf("%d\n", &option)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			fmt.Println("Enter Name of new User: ")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Enter Mobile Number: ")
			var number string
			fmt.Scanln(&number)

			newUser := user.NewUser(name, number)
			fmt.Println("New User Created: ")

			lms.AddUser(newUser)
			clearScreen()

		case 2:
			fmt.Println("Enter Name of new Book: ")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Enter Quantity: ")
			var quantity int
			_, err := fmt.Scanf("%d\n", &quantity)
			if err != nil {
				fmt.Println("Invalid quantity. Please enter a valid number.")
				continue
			}
			newBook := book.NewBook(name, quantity)
			lms.AddBook(newBook)
			clearScreen()

		case 3:
			lms.ViewBooks()

		case 4:
			lms.ViewUsers()

		case 5:
			fmt.Println("Enter Name of Book: ")
			var bookName string
			fmt.Scanln(&bookName)

			fmt.Println("Enter User Mobile Number: ")
			var userNumber string
			fmt.Scanln(&userNumber)

			lms.BorrowBook(userNumber, bookName)
			clearScreen()

		case 6:
			fmt.Println("For Returning a Book, please provide:")
			fmt.Println("Enter Name of Book: ")
			var bookName string
			fmt.Scanln(&bookName)

			fmt.Println("Enter User Mobile Number: ")
			var userNumber string
			fmt.Scanln(&userNumber)

			lms.ReturnBook(userNumber, bookName)
			clearScreen()

		case 7:
			fmt.Println("Enter Number Mobile of User: ")
			var userNumber string
			fmt.Scanln(&userNumber)
			lms.DisplayBorrowedBooks(userNumber)

		case 8:
			fmt.Println("Exiting the application...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
	time.Sleep(1 * time.Second)
}
