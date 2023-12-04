package main

import (
	"fmt"
	"io/ioutil"

	"go.uber.org/zap"
)

func loadInput(sugar *zap.SugaredLogger) string {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		sugar.Errorw("failed to read file",
			"filename", "input",
			"error", err,
		)
	}
	return string(content)
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	fmt.Println(loadInput(sugar))
}
