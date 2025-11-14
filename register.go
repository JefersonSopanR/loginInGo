package main
import (
	"fmt"
	"bufio"
)

func register(r *bufio.Reader, db *db) {
	fmt.Println("   REGISTER  ")
	username, _ := getInput("Enter username: ", r)
	password, _ := getInput("Enter password: ", r)
	if username == "" || password == "" {
		return
	}
	user := &user{
		username: username,
		password: password,
		gmail: "",
		displayName: "",
	}
	db.users = append(db.users, user)
	login(r, db)
	return
}
