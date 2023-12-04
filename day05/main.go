package main

import (
	"os"
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
	sum := 0

	return sum
}

func part2(input string, sugar *zap.SugaredLogger) int {
	lines := strings.Split(input, "\n")
	sum := 0

	return sum
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	input := loadInput(sugar)
	sugar.Infof("%v", input)

	output := day(input, sugar)
	sugar.Infof("%v", output)
}
