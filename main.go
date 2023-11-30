package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model

	Title string
	Author string
	PageCount int

	CheckedOut bool
	CheckedOutDate time.Time
}

const LIBRARY_NAME string = "Go Library"

// main function
func main() {

	// connect to database
	db := connectDB()
	result := db.Find(&Book{})
	fmt.Println(result)

	// initial welcome
	welcome()

	// loop until user exits
	for {

		// get command
		command := getCommand()

		// process the command
		processCommand(command, db)
	}

}

// connect to database function
func connectDB() *gorm.DB {

	// attempt to connect to in-memory sqlite database
	db, err := gorm.Open(sqlite.Open("go_library.sqlite"), &gorm.Config{})

	// check for errors
	if err != nil {
		panic("failed to connect database") // panic stops execution of program
	}

	// auto migrate database
	db.AutoMigrate(&Book{})

	// return database object
	return db
}

// welcome function
func welcome() {
	fmt.Printf("\nWelcome to %s!", LIBRARY_NAME)
	fmt.Println("\nIn this library, you can view, add, edit, and delete books.")
	fmt.Println("\nWe hope you enjoy your stay!")
	fmt.Println("\nType 'help' for a list of commands.")
	fmt.Println("")
}

// get user command function
func getCommand() string {
	var command string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(">> ")
	scanner.Scan()
	command = scanner.Text()

	return command
}

// get book information
func getBookInfo() (string, string, int) {

	// init vars
	var title string
	var author string
	var page_count int

	// create scanner object (needed for reading multiple words)
	scanner := bufio.NewScanner(os.Stdin)

	// get book info
	fmt.Print("\nTitle: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Author: ")
	scanner.Scan()
	author = scanner.Text()

	fmt.Print("Page Count: ")
	fmt.Scan(&page_count)

	fmt.Println("")

	// return book info
	return title, author, page_count
}
