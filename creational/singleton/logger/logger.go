package logger

import (
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	ERROR
	CRITICAL
)

type Logger struct {
	logger   *log.Logger
	logLevel LogLevel
	prefix   string
}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			logger:   log.New(os.Stdout, "[INFO] ", log.LstdFlags),
			logLevel: INFO,
			prefix:   "[INFO] ",
		}
	})
	return instance
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *Logger) updatePrefix() {
	switch l.logLevel {
	case INFO:
		l.prefix = "[INFO] "
	case DEBUG:
		l.prefix = "[DEBUG] "
	case ERROR:
		l.prefix = "[ERROR] "
	case CRITICAL:
		l.prefix = "[CRITICAL] "
	}
	l.logger.SetPrefix(l.prefix)
}

func (l *Logger) log(level LogLevel, message string) {
	l.SetLogLevel(level)
	l.updatePrefix()
	if level >= l.logLevel {
		l.logger.Printf("%s", message)
	}
}

func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *Logger) Critical(message string) {
	l.log(CRITICAL, message)
}
