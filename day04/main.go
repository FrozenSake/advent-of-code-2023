package main

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

func loadInput(sugar *zap.SugaredLogger) string {
	content, err := os.ReadFile("input")
	//content, err := os.ReadFile("example")
	if err != nil {
		sugar.Errorw("failed to read file",
			"filename", "input",
			"error", err,
		)
	}
	return string(content)
}

func part1(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if len(line) <= 1 {
			break
		}
		// Drop Game #
		line = strings.Split(line, ":")[1]
		points := 0

		numbers := strings.Split(line, "|")
		winners := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		for _, winner := range winners {
			if len(winner) == 1 {
				winner = fmt.Sprintf(" %v", winner)
			}
			for _, number := range myNumbers {
				if len(number) == 1 {
					number = fmt.Sprintf(" %v", number)
				}
				if strings.Contains(number, winner) {
					if points >= 1 {
						points = points * 2
					} else {
						points = 1
					}
				}
			}
		}
		sum += points
	}

	return sum
}

func part2(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	var sum int

	// Will use this as a list of how many copies of the upcoming cards are owed, e.g. [1, 2, 3] would be 1 of n+1, 2 of n+2, 3 of n+3, pop off first, add, and go.
	cardCopies := make([]int, len(lines))

	for index, line := range lines {
		if len(line) <= 1 {
			break
		}

		// Add the "base"" card to our copies, then set that number as our win multiplier
		cardCopies[index] += 1
		multiplier := cardCopies[index]

		// Drop Game #
		line = strings.Split(line, ":")[1]
		points := 0

		numbers := strings.Split(line, "|")
		winners := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		for _, winner := range winners {
			if len(winner) == 1 {
				winner = fmt.Sprintf(" %v", winner)
			}
			for _, number := range myNumbers {
				if len(number) == 1 {
					number = fmt.Sprintf(" %v", number)
				}
				if strings.Contains(number, winner) {
					points++
					sugar.Debugf("Card %v won %v times.", index, points)
				}
			}
		}

		sugar.Debugf("Cardcopies length: %v", len(cardCopies))
		for i := 1; i <= points; i++ {
			cardCopies[index+i] += 1 * multiplier
		}
	}

	sugar.Infof("%v", cardCopies)
	for _, value := range cardCopies {
		sum += value
	}

	return sum
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)
	sugar.Infof("%v", input)

	output := part1(input, sugar)
	sugar.Infof("PART ONE: %v", output)

	output64 := part2(input, sugar)
	sugar.Infof("PART 2: %v", output64)
}
