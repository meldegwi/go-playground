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
	l.logf(LevelDebug, DEBUG, format, args...)
}

// Infof formats and prints a message if a log level is info or beyond.
func (l *Logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, INFO, format, args...)
}

// Warnf formats and prints a message if a log level is warn or beyond.
func (l *Logger) Warnf(format string, args ...any) {
	l.logf(LevelWarn, WARN, format, args...)
}

// Errorf formats and prints a message if a log level is error or beyond.
func (l *Logger) Errorf(format string, args ...any) {
	l.logf(LevelError, ERROR, format, args...)
}

// Fatalf formats and prints a message if a log level is fatal.
func (l *Logger) Fatalf(format string, args ...any) {
	l.logf(LevelFatal, FATAL, format, args...)
	os.Exit(1)
}

func (l *Logger) logf(level Level, prefix, format string, args ...any) {
	if l.threshold > level {
		return
	}

	if l.output == nil {
		l.output = os.Stdout
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.output, "%s %s\n", prefix, msg)
}
