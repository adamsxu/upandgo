package upandgo

import (
	"fmt"
	"time"
)

import (
	"log"
	"os"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var level = LevelTrace

func Level() int {
	return level
}

func SetLevel(l int) {
	level = l
}

var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func SetLogger(l *log.Logger) {
	Logger = l
}

func Trace(v ...interface{}) {
	if level <= LevelTrace {
		Logger.Printf("[T] %v\n", v)
	}
}

func Debug(v ...interface{}) {
	if level <= LevelDebug {
		Logger.Printf("[D] %v\n", v)
	}
}

func Info(v ...interface{}) {
	if level <= LevelInfo {
		Logger.Printf("[I] %v\n", v)
	}
}

func Warn(v ...interface{}) {
	if level <= LevelWarning {
		Logger.Printf("[W] %v\n", v)
	}
}

func Error(v ...interface{}) {
	if level <= LevelError {
		Logger.Printf("[E] %v\n", v)
	}
}

func Critical(v ...interface{}) {
	if level <= LevelCritical {
		Logger.Printf("[C] %v\n", v)
	}
}
