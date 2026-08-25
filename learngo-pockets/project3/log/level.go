package log

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used
	// for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information
	// deemed valuable.
	LevelInfo
	// LevelWarn represents a cautious notes about events that
	// unexpected/undesireable but the system recovered from them.
	LevelWarn
	// LevelError represents a high logging level, only to be
	// used to trace errors.
	LevelError
	// LevelFatal represents the highest logging level, only to be
	// used in major error that are unrecoverable.
	LevelFatal
)

// adds prefix indicating its level to the log
const (
	pfxDEBUG = "[DEBUG]"
	pfxINFO  = "[INFO]"
	pfxWARN  = "[WARN]"
	pfxERROR = "[ERROR]"
	pfxFATAL = "[FATAL]"
)

// adds color to the log
const (
	cReset  = "\033[0m"  // no color - rest of the log
	cGray   = "\033[90m" // Debug & timestamp
	cCyan   = "\033[36m" // info
	cYellow = "\033[33m" // warn
	cRed    = "\033[31m" // error
	cPurple = "\033[35m" // fatal
)
