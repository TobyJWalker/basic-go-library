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
		title: book_title, 
		author: book_author, 
		page_count: book_pages,
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