package service

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// The directory for logs to be output
	LogDir = filepath.Join(os.TempDir(), "micro", "logs")
	// The source directory where code lives
	SourceDir = filepath.Join(os.TempDir(), "micro", "uploads")
)

// make the directory
func logFile(serviceName string) string {
	// make the directory
	name := strings.Replace(serviceName, "/", "-", -1)
	return filepath.Join(LogDir, name)
}
