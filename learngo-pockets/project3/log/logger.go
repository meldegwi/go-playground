package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
	msgMxSize int
	colorful  bool
}

// New returns news you a logger, ready to log at the required threshold.
// the default output is Stdout
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout, msgMxSize: 1000, colorful: false}
	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}

// Debugf formats and prints a message if a log level is debug or beyond.
func (l *Logger) Debugf(format string, args ...any) {
	l.logf(LevelDebug, pfxDEBUG, format, args...)
}

// Infof formats and prints a message if a log level is info or beyond.
func (l *Logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, pfxINFO, format, args...)
}

// Warnf formats and prints a message if a log level is warn or beyond.
func (l *Logger) Warnf(format string, args ...any) {
	l.logf(LevelWarn, pfxWARN, format, args...)
}

// Errorf formats and prints a message if a log level is error or beyond.
func (l *Logger) Errorf(format string, args ...any) {
	l.logf(LevelError, pfxERROR, format, args...)
}

// Fatalf formats and prints a message if a log level is fatal.
func (l *Logger) Fatalf(format string, args ...any) {
	l.logf(LevelFatal, pfxFATAL, format, args...)
	os.Exit(1)
}

func (l *Logger) logf(level Level, lvlPfx, format string, args ...any) {
	if l.threshold > level {
		return
	}

	if l.output == nil {
		l.output = os.Stdout
	}

	ts := time.Now().Format("2006-01-02 15:04:05")

	var prefix string
	if l.colorful {
		prefix = addColor(level, ts, lvlPfx)
	} else {
		prefix = fmt.Sprintf("[%v] %s", ts, lvlPfx)
	}

	msg := fmt.Sprintf(format, args...)

	if len(msg) > l.msgMxSize {
		msg = msg[:len(msg)-l.msgMxSize]
	}

	_, _ = fmt.Fprintf(l.output, "%s %s\n", prefix, msg)
}

func addColor(lvl Level, ts, lvlPfx string) string {
	var lvlColor string
	switch lvl {
	case LevelInfo:
		lvlColor = cCyan
	case LevelWarn:
		lvlColor = cYellow
	case LevelError:
		lvlColor = cRed
	case LevelFatal:
		lvlColor = cPurple
	default:
		lvlColor = cGray
	}
	return fmt.Sprintf("%s[%s]%s %s%s%s", cGray, ts, cReset, lvlColor, lvlPfx, cReset)
}
