package utils

import (
	"io"
	"log"
)

type LogLevel uint8

const (
	ERROR   LogLevel = 1 << iota
	WARNING LogLevel = (1 << iota) | ERROR
	INFO    LogLevel = (1 << iota) | WARNING
	DEBUG   LogLevel = (1 << iota) | INFO
)

type FormatedOutType string

const (
	ResetFormatedOut  FormatedOutType = "\x1b[0m"
	RedFormatedOut    FormatedOutType = "\x1b[31m"
	GreenFormatedOut  FormatedOutType = "\x1b[32m"
	YellowFormatedOut FormatedOutType = "\x1b[33m"
	BlueFormatedOut   FormatedOutType = "\x1b[34m"
)

type Logger struct {
	Level LogLevel
	Logs  map[string]*log.Logger
}

func NewLogger(level LogLevel, out io.Writer, flag int) *Logger {
	logLevel := level

	logs := map[string]*log.Logger{
		"ERROR":   log.New(out, string(RedFormatedOut)+"[E ", flag),
		"WARNING": log.New(out, string(YellowFormatedOut)+"[W ", flag),
		"INFO":    log.New(out, string(GreenFormatedOut)+"[I ", flag),
		"DEBUG":   log.New(out, string(BlueFormatedOut)+"[D ", flag),
	}

	return &Logger{
		Level: logLevel,
		Logs:  logs,
	}
}

func addResetSequence(content string) string {
	return "\b] " + string(ResetFormatedOut) + content
}

func (l *Logger) Debug(format string, variables ...interface{}) {
	if (l.Level & DEBUG) == DEBUG {
		format = addResetSequence(format)

		l.Logs["DEBUG"].Printf(format, variables...)
	}
}

func (l *Logger) Info(format string, variables ...interface{}) {
	if (l.Level & INFO) == INFO {
		format = addResetSequence(format)

		l.Logs["INFO"].Printf(format, variables...)
	}
}

func (l *Logger) Warning(format string, variables ...interface{}) {
	if (l.Level & WARNING) == WARNING {
		format = addResetSequence(format)

		l.Logs["WARNING"].Printf(format, variables...)
	}
}

func (l *Logger) Error(format string, variables ...interface{}) {
	if (l.Level & ERROR) == ERROR {
		format = addResetSequence(format)

		l.Logs["ERROR"].Printf(format, variables...)
	}
}
