package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(input string) string {
	return ""
}

func partTwo(input string) string {
	return ""
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading input.txt:")
		fmt.Println(err)
		os.Exit(1)
	}

	input := strings.Trim(string(inputBytes), "\n")
	fmt.Printf("Part One Solution: %s\n", partOne(input))
	fmt.Printf("Part Two Solution: %s\n", partTwo(input))
}
