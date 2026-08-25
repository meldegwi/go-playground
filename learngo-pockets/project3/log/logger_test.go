package log_test

import "logger/log"

func ExampleLogger_Debugf() {
	dl := log.New(log.LevelDebug)
	dl.Debugf("Hello, %s", "world!")
	// Output: [DEBUG] Hello, world!
}
