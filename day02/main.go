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

func day2(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {

		game := strings.Split(lines[i], ":")

		if len(strings.Split(game[0], " ")) < 2 {
			break
		}

		gameNum := strings.Split(game[0], " ")[1]

		redMax := 0
		greenMax := 0
		blueMax := 0

		rounds := strings.Split(game[1], ";")

		for j := 0; j < len(rounds); j++ {
			colours := strings.Split(rounds[j], ",")
			for k := 0; k < len(colours); k++ {
				cubes := strings.Split(colours[k], " ")
				sugar.Debugf("Cubes: %v", cubes)
				switch cubes[2] {
				case "red":
					res, err := strconv.Atoi(cubes[1])
					if err != nil {
						sugar.Errorf("Failure to convert string to int: %s", cubes[0])
					}
					if res > redMax {
						redMax = res
					}
				case "green":
					res, err := strconv.Atoi(cubes[1])
					if err != nil {
						sugar.Errorf("Failure to convert string to int: %s", cubes[0])
					}
					if res > greenMax {
						greenMax = res
					}
				case "blue":
					res, err := strconv.Atoi(cubes[1])
					if err != nil {
						sugar.Errorf("Failure to convert string to int: %s", cubes[0])
					}
					if res > blueMax {
						blueMax = res
					}
				}
			}
		}

		if redMax <= 12 && greenMax <= 13 && blueMax <= 14 {
			sugar.Debugf("Valid game: %v\ngameNum: %v\nredMax: %v, greenMax: %v, blueMax: %v\n", game, gameNum, redMax, greenMax, blueMax)
			num, err := strconv.Atoi(gameNum)
			if err != nil {
				sugar.Errorf("Failure to convert string to int: %s", gameNum)
			}
			sum += num
		}

	}

	return sum
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)

	output := day2(input, sugar)

	sugar.Infof("The sum of the valid games is: %v", output)
}
