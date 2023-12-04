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

/*
func part2(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	sum := 0

	return sum
}
*/

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)
	sugar.Infof("%v", input)

	output := part1(input, sugar)
	sugar.Infof("%v", output)
}
