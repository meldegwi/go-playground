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

const (
	pfxDEBUG = "[DEBUG]"
	pfxINFO  = "[INFO]"
	pfxWARN  = "[WARN]"
	pfxERROR = "[ERROR]"
	pfxFATAL = "[FATAL]"
)

// TODO: add colors to printed console logs.
