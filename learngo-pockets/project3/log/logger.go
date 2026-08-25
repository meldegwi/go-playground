package log

import "fmt"

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

// Debugf formats and prints a message if a log level is debug or beyond.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	message := fmt.Sprintf(format, args...)

	fmt.Printf("[DEBUG] %s", message)
}

// Infof formats and prints a message if a log level is info or beyond.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	message := fmt.Sprintf(format, args...)

	fmt.Printf("[INFO] %s", message)
}

// Warnf formats and prints a message if a log level is warn or beyond.
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	message := fmt.Sprintf(format, args...)

	fmt.Printf("[WARN] %s", message)
}

// Errorf formats and prints a message if a log level is error or beyond.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	message := fmt.Sprintf(format, args...)

	fmt.Printf("[ERROR] %s", message)
}

// Fatalf formats and prints a message if a log level is fatal.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	message := fmt.Sprintf(format, args...)

	fmt.Printf("[FATAL] %s", message)
}
