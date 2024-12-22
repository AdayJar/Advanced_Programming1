package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	isBorrowed bool
}
type Library struct {
	Books map[string]Book
}

func (l Library) AddBook(book Book) {
	_, exists := l.Books[book.ID]
	if exists == true {
		fmt.Println("the book already exists")
		return
	}

	l.Books[book.ID] = book
	fmt.Println("The book was successfully added", book.Title)

}

func (l Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if exists == false {
		fmt.Println("the book dont exist")
		return
	}
	if book.isBorrowed == true {
		fmt.Println("Book already borrowed")
	}
	book.isBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed")

}

func (l Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if exists == false {
		fmt.Println("the book dont exist")
		return
	}
	if book.isBorrowed == false {
		fmt.Println("Book already allow")
	}
	book.isBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned")

}
func (l Library) ListBooks() {
	for _, book := range l.Books {
		if !book.isBorrowed {
			fmt.Println(book.ID, book.Title, book.Author)
		}
	}
}
func main() {
	library := Library{Books: make(map[string]Book)}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nEnter a command (Add, Borrow, Return, List, Exit):")
		if !scanner.Scan() {
			break
		}
		command := strings.TrimSpace(scanner.Text())

		switch strings.ToLower(command) {
		case "add":
			fmt.Println("Enter the book ID:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter the book title:")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter the book author:")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			library.AddBook(Book{ID: id, Title: title, Author: author})

		case "borrow":
			fmt.Println("Enter the book ID to borrow:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			library.BorrowBook(id)

		case "return":
			fmt.Println("Enter the book ID to return:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			library.ReturnBook(id)

		case "list":
			library.ListBooks()

		case "exit":
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}
