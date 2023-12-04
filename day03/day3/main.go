package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func loadInput(sugar *zap.SugaredLogger) string {
	content, err := os.ReadFile("input")
	if err != nil {
		sugar.Errorw("failed to read file",
			"filename", "input",
			"error", err,
		)
	}
	return string(content)
}

func parseLine(line string, y int, sugar *zap.SugaredLogger) int {
	sum := 0

	numbers := "0123456789"
	offset := 1

	if len(line) <= 2 {
		return 0
	}

	num := ""
	// Check y-1
	for true {
		if (y - offset) < 0 {
			sugar.Debug("y - offset < 0, cannot continue")
			break
		} else if strings.Contains(numbers, string(line[y-offset])) {
			// new character before num because we're moving right to left.
			num = fmt.Sprintf("%s%s", string(line[y-offset]), num)
		} else {
			sugar.Debug("No more digits on the left.")
			break
		}
		offset++
	}

	// check y
	if strings.Contains(numbers, string(line[y])) {
		// num before new character because if we already have a number, we're moving left to right now.
		num = fmt.Sprintf("%s%s", num, string(line[y]))
	} else {
		if num != "" {
			number, err := strconv.Atoi(num)
			if err != nil {
				sugar.Errorf("Could not convert %s to number: %v", num, err)
			}
			sum += number
			num = ""
		}
	}

	// check y+1
	offset = 1
	for true {
		if (y + offset) >= len(line) {
			sugar.Debug("y + offset is out out of bounds, cannot continue")
			break
		} else if strings.Contains(numbers, string(line[y+offset])) {
			// num before new character because if we already have a number, we're moving left to right now.
			num = fmt.Sprintf("%s%s", num, string(line[y+offset]))
		} else {
			sugar.Debug("No more digits on the right.")
			break
		}
		offset++
	}

	// If we have a number being processed, convert it now
	if num != "" {
		number, err := strconv.Atoi(num)
		if err != nil {
			sugar.Errorf("Could not convert %s to number: %v", num, err)
		}
		sum += number
		num = ""
	}

	sugar.Debugf("Called on line: %v\nSymbol pos: %v\nSum for symbol: %v\n", line, y, sum)

	return sum
}

func day3(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	sum := 0

	symbols := "#%&*+-/=@$"

	var lastLine string

	// This logic falls apart if there's ever a number adjacent to two symbols because it won't be excluded. Would need list of points and to validate the point wasn't already recorded to fix.
	for i := 0; i < len(lines); i++ {

		// Iterate the string looking for symbols.
		for j, character := range lines[i] {
			char := string(character)
			if strings.Contains(symbols, char) {
				sum += parseLine(lastLine, j, sugar)
				sugar.Debugf("Sum: %v\n", sum)
				sum += parseLine(lines[i], j, sugar)
				sugar.Debugf("Sum: %v\n", sum)
				if i+1 < len(lines) {
					sum += parseLine(lines[i+1], j, sugar)
					sugar.Debugf("Sum: %v\n", sum)
				}
				sugar.Debugf("------")
			}
		}
		lastLine = lines[i]
	}

	return sum
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)
	sugar.Infof("%v", input)

	output := day3(input, sugar)
	sugar.Infof("%v", output)
}
