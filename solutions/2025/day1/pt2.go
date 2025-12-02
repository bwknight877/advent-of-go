package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

// submit: go run . -s -y 2025 -d 1 -p 2

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	var count int
	var err error

	// start position
	position := 50

	for scanner.Scan() {
		line := scanner.Text()

		if err = turnDial2(&position, &count, line); err != nil {
			return "", fmt.Errorf("turn dial from position %v line %q: %w", position, line, err)
		}
	}

	return strconv.Itoa(count), nil
}

func turnDial2(position *int, count *int, line string) error {
	if len(line) < 2 {
		return fmt.Errorf("invalid line %q", line)
	}

	clicks, err := strconv.Atoi(line[1:])
	if err != nil {
		return fmt.Errorf("atoi %q: %w", line[1:], err)
	}

	newCount := *count

	// count for each multiple of 100
	if clicks > 99 {
		var fullRotations int = clicks / 100

		// add the full rotations to the count
		newCount += fullRotations

		// adjust the click count without the full rotations
		clicks %= 100
	}

	switch line[0] {
	case 'R':
		// positive offset

	case 'L':
		// negative offset
		clicks *= -1

	default:
		return fmt.Errorf("invalid value %v", line[0])
	}

	// new position is the current position plus the click offset
	newPosition := *position + clicks

	if newPosition > 99 {
		// if we have gone past 99, count a click
		newPosition -= 100
		newCount += 1
	} else if newPosition == 0 {
		// if we are at 0, count a click
		newCount += 1
	} else if newPosition < 0 && *position != 0 {
		// if we are below 0, and we didn't start at 0, count a click
		newPosition += 100
		newCount += 1
	}

	fmt.Println("position", *position, "count", *count, "line", line, "clicks", clicks, "newPosition", newPosition, "newCount", newCount)

	if newPosition < 0 {
		// add 100 if negative
		newPosition += 100
	}

	*position = newPosition
	*count = newCount

	return nil
}
