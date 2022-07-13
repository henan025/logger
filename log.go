package logger

import (
	"fmt"
)

type Level int

const (
	VERBOSE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

func (l Level) String() string {
	switch l {
	case VERBOSE:
		return "Verbose"
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case WARNING:
		return "Warning"
	case ERROR:
		return "Error"
	}
	return "unknown"
}

type Output interface {
	Init(opts *LogOptions)
	Output(msg string)
}

type LogOptions struct {
	level      Level
	adapter    string
	outputPath string
}

type Logger struct {
	level  Level
	output Output
}

func (logger *Logger) Initialize(opts *LogOptions) {
	if opts == nil {
		logger.level = INFO
		logger.output = &Stdout{}
	} else {
		logger.level = opts.level
		if opts.adapter == "file" {
			logger.output = &FileOutput{}
		} else {
			logger.output = &Stdout{}
		}
	}
	logger.output.Init(opts)
}

func (logger *Logger) log(level Level, txt string) {
	logger.output.Output(fmt.Sprintf("[%s]: %s", level.String(), txt))
}

func (logger *Logger) logln(level Level, v ...interface{}) {
	logger.log(level, fmt.Sprintln(v...))
}

func (logger *Logger) logf(level Level, format string, v ...interface{}) {
	logger.log(level, fmt.Sprintf(format, v...))
}

func (logger *Logger) Verbose(v ...interface{}) {
	logger.logln(VERBOSE, v...)
}

func (logger *Logger) Verbosef(format string, v ...interface{}) {
	logger.logf(VERBOSE, format, v...)
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.logln(DEBUG, v...)
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.logf(DEBUG, format, v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.logln(INFO, v...)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.logf(INFO, format, v...)
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.logln(WARNING, v...)
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.logf(WARNING, format, v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.logln(ERROR, v...)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.logf(ERROR, format, v...)
}
