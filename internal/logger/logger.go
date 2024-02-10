package logger

import (
	"os"
	"path"
	"syscall"

	"github.com/rs/zerolog"
)

type (
	Logger interface {
		// Get logger instance
		GetInstance() any

		//Log error
		Error(message string)

		//Log error
		Info(message string)

		//Log error
		Debug(message string)
	}

	logger struct {
		instance zerolog.Logger
	}

	fileLoggerOutput struct {
		filePath string
	}
)

func (f *fileLoggerOutput) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(f.filePath, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	return file.Write(p)
}

func newFileLoggerOutput(path string) *fileLoggerOutput {
	return &fileLoggerOutput{
		filePath: path,
	}
}

func NewLogger() Logger {
	instance := zerolog.New(newFileLoggerOutput(path.Join(".", "./logs.log")))

	return &logger{
		instance: instance,
	}
}

func (l *logger) GetInstance() any {
	return l.instance
}

func (l *logger) Error(message string) {
	l.instance.Error().Msg(message)
}

func (l *logger) Info(message string) {
	l.instance.Info().Msg(message)
}

func (l *logger) Debug(message string) {
	l.instance.Debug().Msg(message)
}
