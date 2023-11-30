package main

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

// process command
func processCommand(command string, db *gorm.DB) {

	// parse command
	parsed_command := parseCommand(command)
	action := parsed_command[0]

	// switch on action
	switch action {

	case "help":
		help()

	case "exit":
		exit()

	case "add":
		addBook(db)

	case "list":
		listBooks(db)

	case "delete":
		// check for book name
		if len(parsed_command) < 2 {
			// print error message
			fmt.Println("Missing book name. Try 'delete <book>'.")
		} else if parsed_command[2] != "" {
			// print error message
			fmt.Println("Too many arguments. Try 'delete <book>'.")
		} else {
			// delete book
			deleteBook(db, parsed_command[1])
		}
	
	case "view":
		// check for book name
		if len(parsed_command) < 2 {
			// print error message
			fmt.Println("Missing book name. Try 'view <book>'.")
		} else if parsed_command[2] != "" {
			// print error message
			fmt.Println("Too many arguments. Try 'view <book>'.")
		} else {
			// view book
			viewBook(db, parsed_command[1])
		}

	default:
		// print error message
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
	}
}

// parse a command
func parseCommand(command string) [3]string {

	// array to hold command words
	var command_words [3]string

	// split command into words
	split_command := strings.Split(command, " ")
	
	// copy words into array
	copy(command_words[:], split_command)

	// return array
	return command_words
}

// help function
func help() {
	
	// print help message
	fmt.Println("\nAvailable commands:")
	fmt.Println("\nhelp - display this message")
	fmt.Println("list - list all books")
	fmt.Println("add - add a book")
	fmt.Println("view <book> - view book information")
	fmt.Println("delete <book> - delete a book")
	fmt.Println("checkout - checkout a book")
	fmt.Println("checkin - checkin a book")
	fmt.Println("exit - exit the program")
	fmt.Println("")
}

// exit function
func exit() {

	// exit msg
	fmt.Println("\nThank you for using the library!")
	fmt.Println("")

	// exit program
	os.Exit(0)
}

// add book function
func addBook(db *gorm.DB) {

	// init vars
	var book_title string
	var book_author string
	var book_pages int

	// get book info
	book_title, book_author, book_pages = getBookInfo()

	// create book object
	book := Book{
		Title: book_title, 
		Author: book_author, 
		PageCount: book_pages,
		CheckedOut: false,
	}

	// add book to database
	result := db.Create(&book)

	// print success message
	if result.Error == nil {
		fmt.Println("\nBook added successfully!")
		fmt.Println("")
	} else {
		fmt.Println("\nError adding book.")
		fmt.Println("")
	}

}

// list books function
func listBooks(db *gorm.DB) {
	
	// init book array
	var books []Book

	// get books from database
	db.Find(&books)

	// print books
	fmt.Println("\nBooks:")
	for i, book := range books {
		fmt.Printf("%d: %s\n", i + 1, book.Title)
	}
	fmt.Println("")
}

// delete book function
func deleteBook(db *gorm.DB, book_name string) {

	// init book object
	var book Book

	// get book from database
	result := db.Where("title = ?", book_name).First(&book)

	// check for errors
	if result.Error != nil {
		// print error message
		fmt.Println("\nUnable to find book.")
		fmt.Println("")
		return
	}

	// delete book from database
	result = db.Delete(&book)

	// check for errors
	if result.Error == nil {
		// print success message
		fmt.Println("\nBook deleted successfully!")
		fmt.Println("")
	} else {
		// print error message
		fmt.Println("\nUnable to delete book.")
		fmt.Println("")
	}
}

// view book function
func viewBook(db *gorm.DB, book_name string) {

	// init book object
	var book Book

	// get book from database
	result := db.Where("title = ?", book_name).First(&book)

	// check for errors
	if result.Error != nil {
		// print error message
		fmt.Println("\nUnable to find book.")
		fmt.Println("")
		return
	}

	// print book info
	fmt.Println("\nBook:")
	fmt.Printf("Title: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Pages: %d\n", book.PageCount)
	
	// print checked out status
	if book.CheckedOut {
		fmt.Println("Checked Out: Yes")
		fmt.Printf("Checked Out Date: %s\n", book.CheckedOutDate)
	} else {
		fmt.Println("Checked Out: No")
	}

	fmt.Println("")

}