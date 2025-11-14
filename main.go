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

	register(reader, &db)
}
