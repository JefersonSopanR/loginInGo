package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	db := db{}
	reader := bufio.NewReader(os.Stdin)

	input, _ := getInput("Enter the fileName: ", reader)
	for input != "exit" {
		if input == "register" {
			register(reader, &db)
		} else if input == "login" {
			login(reader, &db)
		}
		input, _ = getInput("Enter the fileName: ", reader)
	}
}
