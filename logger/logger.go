package logger

import (
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	PANIC
	FATAL
)

var (
	loggers map[Level]*log.Logger
)

func init() {
	loggers = make(map[Level]*log.Logger)
	loggers[DEBUG] = log.New(os.Stdout, "[DEBUG]", log.Ltime)
	loggers[INFO] = log.New(os.Stdout, "[INFO]", log.Ltime)
	loggers[WARNING] = log.New(os.Stderr, "[WARNING]", log.Ltime)
	loggers[ERROR] = log.New(os.Stderr, "[ERROR]", log.Ltime)
	loggers[PANIC] = log.New(os.Stderr, "[PANIC]", log.Ltime)
	loggers[FATAL] = log.New(os.Stderr, "[FATAL]", log.Ltime)
}

func Log(level Level, v ...any) {
	switch level {
	case DEBUG, INFO, WARNING, ERROR:
		loggers[level].Println(v...)
	case PANIC:
		loggers[level].Panicln(v...)
	case FATAL:
		loggers[level].Fatalln(v...)
	}
}

func Logf(level Level, format string, v ...any) {
	switch level {
	case DEBUG, INFO, WARNING, ERROR:
		loggers[level].Printf(format, v...)
	case PANIC:
		loggers[level].Panicf(format, v...)
	case FATAL:
		loggers[level].Fatalf(format, v...)
	}
}

func Debug(v ...any) {
	Log(DEBUG, v...)
}

func Info(v ...any) {
	Log(INFO, v...)
}

func Warning(v ...any) {
	Log(WARNING, v...)
}

func Error(v ...any) {
	Log(ERROR, v...)
}

func Panic(v ...any) {
	Log(PANIC, v...)
}

func Fatal(v ...any) {
	Log(FATAL, v...)
}

func Debugf(format string, v ...any) {
	Logf(DEBUG, format, v...)
}

func Infof(format string, v ...any) {
	Logf(INFO, format, v...)
}

func Warningf(format string, v ...any) {
	Logf(WARNING, format, v...)
}

func Errorf(format string, v ...any) {
	Logf(ERROR, format, v...)
}

func Panicf(format string, v ...any) {
	Logf(PANIC, format, v...)
}

func Fatalf(format string, v ...any) {
	Logf(FATAL, format, v...)
}
