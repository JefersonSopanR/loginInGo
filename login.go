package main

import (
	"fmt"
	"bufio"
)

type form struct {
	username string
	password string
}



func login(r *bufio.Reader, db *db) {
	fmt.Println("   LOGIN   ")
	promptUsername := "Enter the username: "
	promptPassword := "Enter the password: "
	username, _ := getInput(promptUsername, r)
	password, _ := getInput(promptPassword, r)
	if username == "" || password == "" {
		fmt.Println("Invalid Username or Password")
		login(r, db)
		return
	}
	f := form{
		username: username,
		password: password,
	}
	user := goDataBase(f, db)
	if user != nil {
		cardUser(user, db)
		return
	} else {
		fmt.Println("Invalid Username or Password")
		login(r, db)
		return
	}
}

