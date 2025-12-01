package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func parseRotation(rot string) int {
	amount, _ := strconv.Atoi(rot[1:])

	if (rot[0] == 'L') {
		return -amount
	}

	return amount
}

func partOne(input string) string {
	rotations := strings.Split(input, "\n")
	pos := 50
	zeroCount := 0

	for _, rot := range rotations {
		pos += parseRotation(rot)

		pos %= 100
		if (pos < 0) {
			pos += 100
		}

		if (pos == 0) {
			zeroCount++
		}
	}

	return fmt.Sprint(zeroCount)
}

func partTwo(input string) string {
	rotations := strings.Split(input, "\n")
	pos := 50
	zeroCount := 0

	for _, rot := range rotations {
		amount := parseRotation(rot)
		pos += amount

		switch {
		case pos == 0:
			zeroCount++
		case pos > 0:
			zeroCount += pos / 100
		case pos == amount:
			// Started at zero and went negative, only passed zero if less than -100
			zeroCount += -pos / 100
		default:
			// Started above zero and went negative, passed zero at least once
			zeroCount += -pos / 100 + 1
		}

		pos %= 100
		if (pos < 0) {
			pos += 100
		}
	}

	return fmt.Sprint(zeroCount)
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
