package logging

import (
	"errors"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	JsonFormat = "json"
	TextFormat = "text"
)

func Setup(logFormat string) error {

	switch logFormat {
	case TextFormat:
		// in text mode we use FormatSerialize format
		log.SetLogger(klogr.New())
		return nil
	case JsonFormat:
		zapLog, err := zap.NewProduction()
		if err != nil {
			return ErrZapError

		}
		klog.SetLogger(zapr.NewLogger(zapLog))

		// in json mode we use FormatKlog format
		log.SetLogger(klog.NewKlogr())
		return nil
	}
	return ErrorFormatError
}

//Returns the global logger as configured in main
func globalLogger() logr.Logger {
	return log.Log
}

//Return a new Logger instance with the specified name element added to the Logger's name
func WithName(name string) logr.Logger {
	return globalLogger().WithName(name)
}

//WithValues returns a new Logger instance with with additional key/value pairs
func WithValues(keysAndValues ...interface{}) logr.Logger {
	return globalLogger().WithValues(keysAndValues...)
}

//V returns a new Logger instance for a specific verbosity level
func V(level int) logr.Logger {
	return globalLogger().V(level)
}

func Info(msg string, kvs ...interface{}) {
	globalLogger().Info(msg, kvs...)
}

var ErrZapError = errors.New("JSON logger could not be initialized")

var ErrorFormatError = errors.New("log format not recognized, pass `text` for text mode or `json` to enable JSON logging")
