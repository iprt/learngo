package main

import "go.uber.org/zap"

func main() {

	logger, _ := zap.NewProduction()

	logger.Debug("this is an debug message")

	logger.Error("this is an error message")
	logger.Error("this is an error message")
	logger.Error("this is an error message")
	logger.Error("this is an error message")
	logger.Error("this is an error message")
}
