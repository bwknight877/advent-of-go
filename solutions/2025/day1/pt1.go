package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

// submit: go run . -s -y 2025 -d 1 -p 1

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	var count int
	var err error

	// start position
	position := 50

	for scanner.Scan() {
		line := scanner.Text()

		position, err = turnDial(position, line)
		if err != nil {
			return "", fmt.Errorf("turn dial from position %v line %q: %w", position, line, err)
		}

		// increment count if position is 0
		if position == 0 {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func turnDial(position int, line string) (int, error) {
	if len(line) < 2 {
		return 0, fmt.Errorf("invalid line %q", line)
	}

	clicks, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, fmt.Errorf("atoi %q: %w", line[1:], err)
	}

	switch line[0] {
	case 'R':
		// positive offset

	case 'L':
		// negative offset
		clicks *= -1

	default:
		return 0, fmt.Errorf("invalid value %v", line[0])
	}

	// combine the position and the number parsed from the line
	newPosition := position + clicks

	// always mod 100
	newPosition %= 100

	if newPosition < 0 {
		// add 100 if negative
		newPosition += 100
	}

	return newPosition, nil
}
