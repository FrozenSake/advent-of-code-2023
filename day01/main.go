package main

import (
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

func part1(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")

	sumOfCalibration := 0

	for i := 0; i < len(lines); i++ {
		loopNum := 0
		firstNum := -1
		secondNum := -1
		for j := 0; j < len(lines[i]); j++ {
			if firstNum < 0 {
				character := string(lines[i][loopNum])
				if result, err := strconv.Atoi(character); err == nil {
					firstNum = result
				}
			}
			if secondNum < 0 {
				character := string(lines[i][len(lines[i])-1-loopNum])
				if result, err := strconv.Atoi(character); err == nil {
					secondNum = result
				}
			}
			if firstNum >= 0 && secondNum >= 0 {
				sumOfCalibration = sumOfCalibration + (firstNum*10 + secondNum)
				sugar.Infof("On line: %v, and Loop: %v\nfirstNum: %v, secondNum: %v\nsumOfCalibration: %v", i, loopNum, firstNum, secondNum, sumOfCalibration)
				break
			}
			loopNum++
		}
	}

	return sumOfCalibration
}

func part2(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")

	sumOfCalibration := 0

	return sumOfCalibration
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)

	output := day1(input, sugar)

	sugar.Infof("The sum of calibrations for day 1 is: %v\n", output)
}
