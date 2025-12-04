// https://adventofcode.com/2025/day/3

package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func findMaxDigitIndex(numberString string) int {
	index := -1
	largest := -1

	for i, char := range numberString {
		digit, err := strconv.Atoi(string(char))
		if digit > largest && err == nil {
			largest = digit
			index = i
		}
	}

	return index
}

func findLargestJoltage(battery string) int {
	firstIdx := findMaxDigitIndex(battery[:len(battery) - 1])
	offsetSecondIdx := findMaxDigitIndex(battery[firstIdx + 1:])

	firstDigit := string(battery[firstIdx]);
	secondDigit := string(battery[firstIdx + offsetSecondIdx + 1])

	joltage, _ := strconv.Atoi(firstDigit + secondDigit)
	return joltage
}

func findLargestJoltageOf(length int, battery string) int {
	joltage := ""
	pos := 0

	for i := 0; i < length; i++ {
		valid := battery[pos: len(battery) - (length - i - 1)]
		idx := findMaxDigitIndex(valid)
		joltage += string(valid[idx])
		pos += idx + 1
	}

	asInt, _ := strconv.Atoi(joltage)
	return asInt
}


func partOne(input string) string {
	batteries := strings.Split(input, "\n")
	totalJoltage := 0

	for _, battery := range batteries {
		totalJoltage += findLargestJoltage(battery)
	}

	return fmt.Sprint(totalJoltage)
}

func partTwo(input string) string {
	batteries := strings.Split(input, "\n")
	totalJoltage := 0

	for _, battery := range batteries {
		totalJoltage += findLargestJoltageOf(12, battery)
	}

	return fmt.Sprint(totalJoltage)
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
