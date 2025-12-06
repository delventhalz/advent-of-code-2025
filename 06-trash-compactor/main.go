// https://adventofcode.com/2025/day/6

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type CephalapodMath struct {
	operator byte
	operands []int
}

func splitSpace(str string) []string {
	var split []string
	current := ""

	for _, ch := range str {
		if unicode.IsSpace(ch) {
			if current != "" {
				split = append(split, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		split = append(split, current)
	}

	return split
}

func parseOperandLine(line string) []int {
	tokens := splitSpace(line)
	operands := make([]int, len(tokens))

	for i, asStr := range tokens {
		asInt, _ := strconv.Atoi(asStr)
		operands[i] = asInt
	}

	return operands
}

func normalizeLines(lines []string) {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	for i, line := range lines {
		for j := len(line); j < maxLen; j++ {
			lines[i] += " "
		}
	}
}

func parseCephalapodMath(worksheet string) []CephalapodMath {
	var maths []CephalapodMath
	var current CephalapodMath

	lines := strings.Split(worksheet, "\n")
	normalizeLines(lines)
	opLine := lines[len(lines) - 1]
	operandLines := lines[0: len(lines) - 1]

	for col := 0; col < len(lines[0]); col++ {
		if opLine[col] != ' ' {
			maths = append(maths, current)
			current = CephalapodMath{operator: opLine[col]}
		}

		operand := ""
		for row := 0; row < len(operandLines); row++ {
			operand += string(operandLines[row][col])
		}
		operand = strings.TrimSpace(operand)

		if (operand != "") {
			asInt, err := strconv.Atoi(operand)
			if err != nil {
				fmt.Printf("Failed to parse operand at %d: %q\n", col, operand)
				fmt.Println(err)
				os.Exit(1)
			}

			current.operands = append(current.operands, asInt)
		}
	}

	maths = append(maths, current)

	// First math is an empty starting value
	return maths[1:]
}

func partOne(input string) string {
	lines := strings.Split(input, "\n")
	operators := splitSpace(lines[4])

	operandsByIndex := [4][]int {
		parseOperandLine(lines[0]),
		parseOperandLine(lines[1]),
		parseOperandLine(lines[2]),
		parseOperandLine(lines[3]),
	}

	grandTotal := 0

	for i, start := range operandsByIndex[0] {
		result := start

		for j := 1; j < 4; j++ {
			if (operators[i] == "*") {
				result *= operandsByIndex[j][i]
			} else {
				result += operandsByIndex[j][i]
			}
		}

		grandTotal += result
	}

	return fmt.Sprint(grandTotal)
}

func partTwo(input string) string {
	maths := parseCephalapodMath(input)
	grandTotal := 0

	for _, math := range maths {
		result := math.operands[0]

		for _, operand := range math.operands[1:] {
			if (math.operator == '*') {
				result *= operand
			} else {
				result += operand
			}
		}

		grandTotal += result
	}

	return fmt.Sprint(grandTotal)
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
