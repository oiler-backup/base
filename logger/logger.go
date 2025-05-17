package logger

import "go.uber.org/zap"

type Mode int

const (
	PRODUCTION Mode = iota
	DEVELOPMENT
)

func GetLogger(mode Mode) (*zap.SugaredLogger, error) { // coverage-ignore
	var logger *zap.Logger
	var err error
	if mode == PRODUCTION {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	}
	return logger.Sugar(), nil
}
