package logger

import "go.uber.org/zap"

type Logger struct {
	Logger *zap.Logger
}


func (l *Logger) Init() error {
	return nil
}

func (l *Logger) Configure() error {
	return nil
}

func (l *Logger) Serve() chan error {
	errCh := make(chan error)
	return errCh
}