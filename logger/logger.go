// Package logger provides a convenience function to constructing a logger
// for use. This is required not just for applications but for testing.
package logger

import (
	"github.com/Loner1024/uniix.io/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New constructs a Sugared Logger that writes to stdout and
// provides human-readable timestamps.
func New(conf configs.Config) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": conf.App.Name,
	}
	
	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	
	return log.Sugar(), nil
}