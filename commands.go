package main

import (
	"fmt"
	"strings"
)

// process command
func processCommand(command string) {

	// parse command
	parsed_command := parseCommand(command)
	action := parsed_command[0]

	// switch on action
	switch action {

	case "help":
		help()

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
	fmt.Println("create - add a book")
	fmt.Println("view <book> - view book information")
	fmt.Println("delete <book> - delete a book")
	fmt.Println("checkout - checkout a book")
	fmt.Println("checkin - checkin a book")
	fmt.Println("exit - exit the program")
	fmt.Println("")
}