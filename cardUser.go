package main

import (
	"fmt"
	"bufio"
	"os"
)

func cardUser(u *user, db *db) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nWELCOME %v TO THE MAIN PAGE\n", u.username)
	fmt.Printf("	username: %v\n", u.username)
	if u.gmail != "" {
		fmt.Printf("	email: %v\n")
	}
	if u.displayName != "" {
		fmt.Printf("	DisplayName: %v\n", u.displayName)
	}
	fmt.Printf("To interact type:\nU) to update the user\nE) to logout\nD) to delete user\n")
	input, _ := getInput(" > ", reader)
	switch input {
	case "U":
		fmt.Println(input)
		cardUser(u, db)
		return
	case "E":
		fmt.Println(input)
		login()
		return
	case "D":
		fmt.Println(input)
		cardUser(u, db)
		return
	default:
		fmt.Println("Invalid type please enter a valid one!")
		cardUser(u, db)
		return
	}
	return
}
