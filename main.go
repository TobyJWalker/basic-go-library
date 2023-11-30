package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model

	title string
	author string
	page_count int

	checked_out bool
	checked_out_date time.Time
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
		processCommand(command)
	}

}

// connect to database function
func connectDB() *gorm.DB {

	// attempt to connect to in-memory sqlite database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

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
	fmt.Print(">> ")
	fmt.Scanln(&command)

	return command
}
