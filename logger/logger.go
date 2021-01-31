package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

//func (l *Logger) Message() Message

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}