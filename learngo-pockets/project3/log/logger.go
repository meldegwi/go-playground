// Package log...
package log

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// New returns news you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if a log level is debug.
func (l *Logger) Debugf(format string, args ...any) {
	// implement me.
}

// Infof formats and prints a message if a log level is info.
func (l *Logger) Infof(format string, args ...any) {
	// implement me.
}

// Errorf formats and prints a message if a log level is error.
func (l *Logger) Errorf(format string, args ...any) {
	// implement me.
}
