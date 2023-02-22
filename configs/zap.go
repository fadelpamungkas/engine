package configs

import (
	"go.uber.org/zap"
)

func InitializeZap() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
