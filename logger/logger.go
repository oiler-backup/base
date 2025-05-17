// Package logger provides logger for use in the project.
//
// The returned logger is either Zap Production or Development logger.
package logger

import "go.uber.org/zap"

// A mode is an enum for selecting type of logger.
type Mode int

// The resulting logger type is caclulated according to this modes.
const (
	PRODUCTION Mode = iota
	DEVELOPMENT
)

// GetLogger returns new Zap logger or error if occured.
// Logger instance can be either Production or Development depending on mode.
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
