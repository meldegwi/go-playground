package log

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns news you a logger, ready to log at the required threshold.
// the default output is Stdout
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if a log level is debug or beyond.
func (l *Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		l.logf(format, DEBUG, args...)
	}
}

// Infof formats and prints a message if a log level is info or beyond.
func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelInfo {
		l.logf(format, INFO, args...)
	}
}

// Warnf formats and prints a message if a log level is warn or beyond.
func (l *Logger) Warnf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelWarn {
		l.logf(format, WARN, args...)
	}
}

// Errorf formats and prints a message if a log level is error or beyond.
func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stderr
	}

	if l.threshold <= LevelError {
		l.logf(format, ERROR, args...)
	}
}

// Fatalf formats and prints a message if a log level is fatal.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stderr
	}

	if l.threshold <= LevelFatal {
		l.logf(format, FATAL, args...)
	}
}

func (l *Logger) logf(format string, level string, args ...any) {
	_, _ = fmt.Fprintf(l.output, level+" "+format+"\n", args...)
}
