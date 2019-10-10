package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() error {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		return err
	}
	Logger = logger
	return nil
}
