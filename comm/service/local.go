package service

import (
	"os"
	"path/filepath"
	"strings"
)

// defaultNamespace to use if not provided as an option
const defaultNamespace = "services"

var (
	// The directory for logs to be output
	LogDir = filepath.Join(os.TempDir(), defaultNamespace, "logs")
)

// make the directory
func logFile(serviceName string) string {
	name := strings.Replace(serviceName, "/", "-", -1)
	fp := filepath.Join(LogDir, name)
	return fp
}
