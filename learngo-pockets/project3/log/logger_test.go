package log_test

import (
	"os"

	"logger/log"
)

func ExampleLogger_Debugf() {
	dl := log.New(log.LevelDebug, os.Stdout)
	dl.Debugf("Hello, %s", "world!")
	// Output: [DEBUG] Hello, world!
}
