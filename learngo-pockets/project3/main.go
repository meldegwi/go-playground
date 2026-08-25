package main

import (
	"os"
	"time"

	"logger/log"
)

func main() {
	lgr := log.New(log.LevelInfo, log.WithOutput(os.Stdout))

	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Error are values, documentation is for %s", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)
	lgr.Infof("Hello! the time now is %v", time.Now())
	lgr.Fatalf("Something terribly bad happened :(")
}
