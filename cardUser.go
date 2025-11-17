package main

import (
	"fmt"
	"bufio"
	"os"
)

func cardUser(u *user, db *db) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nWELCOME %v TO THE MAIN PAGE\n", u.username)
	fmt.Printf("	Username: %v\n", u.username)
	if u.gmail != "" {
		fmt.Printf("	Gmail: %v\n", u.gmail)
	}
	if u.displayName != "" {
		fmt.Printf("	DisplayName: %v\n", u.displayName)
	}
	fmt.Printf("To interact type:\nU) to update the user\nL) to logout\nD) to delete user\n")
	input, _ := getInput(" > ", reader)
	switch input {
	case "U":
		fmt.Println(input)
		updateUser(u, reader)
		cardUser(u, db)
		return
	case "L":
		fmt.Println(input)
		login(reader, db)
		return
	case "D":
		fmt.Println(input)
		db.deleteUser(u)
		login(reader, db)
		return
	default:
		fmt.Println("Invalid type please enter a valid one!")
		cardUser(u, db)
		return
	}
	return
}

func updateUser(u *user, r *bufio.Reader) {
	fmt.Println("CHOSSE WHAT YOU TO WANT TO CHANGE!")
	gmail := u.gmail
	if u.gmail == "" {
		gmail = "/none/"
	}
	displayName := u.displayName
	if u.displayName == "" {
		displayName = "/none/"
	}

	fmt.Printf("U) Username: %v\n", u.username)
	fmt.Printf("G) gmail: %v\n", gmail)
	fmt.Printf("D) Display Name: %v\n", displayName)
	fmt.Printf("Q) QUIT\n")

	option, _:= getInput(" > option --> ", r)

	switch option {
	case "U":
		updateUsername(u, r)
	case "G":
		updategmail(u, r)
	case "D":
		updateDisplayName(u, r)
	case "Q":
		return
	default:
		fmt.Println("You chose not a valid option")
		updateUser(u, r)
		return
	}
}

func updateUsername(u *user, r *bufio.Reader) {

	newUsename, err := getInput("Enter new username: ", r)

	if newUsename == "" || err != nil {
		fmt.Println("Not valid username")
		return 
	}
	u.username = newUsename
	fmt.Println("Username updated succesfully!")
	return
}

func updategmail(u *user, r *bufio.Reader) {

	newgmail, err := getInput("Enter new gmail: ", r)

	if newgmail == "" || err != nil {
		fmt.Println("Not valid gmail")
		return 
	}
	u.gmail = newgmail
	fmt.Println("gmail updatedcardUser(u, db) succesfully!")
	return
}

func updateDisplayName(u *user, r *bufio.Reader) {

	newDisplayName, err := getInput("Enter new displayName: ", r)

	if newDisplayName == "" || err != nil {
		fmt.Println("Not valid displayName")
		return 
	}
	u.displayName = newDisplayName
	fmt.Println("DisplayName updated succesfully!")
	return
}

