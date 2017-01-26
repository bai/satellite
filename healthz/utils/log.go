package utils

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/trace"
)

// SetupLogging configures logging: sets up log level, output to stderr and
// formatter TextFormatter
func SetupLogging(level string) error {
	lvl := strings.ToLower(level)
	if lvl == "debug" {
		trace.EnableDebug()
	}
	sev, err := log.ParseLevel(lvl)
	if err != nil {
		return err
	}
	log.SetLevel(sev)

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	return nil
}
