/*
Package log exposes an API to log your work.

First, instantiate a logger with log.New(), and give it a threshold level.

Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on three levels:
  - Debug: mostly used to debug code, follow step by step processes
  - Info: valuable messages providing insights to milestones of a process
  - Warn: warning messages about change in a process that got recovered or similar
  - Error: error message to understand what went wrong.
  - Fatal: message about error that panicked a process/system and is unrecoverable
*/
package log
