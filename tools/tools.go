package tools

import (
	"fmt"
	"os"
	"strings"
)

func GetEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s key not found", key)
	}

	return value, nil
}

func ConvertWindowDirToLinuxDir(path string) string {
	split := strings.Split(path, string(os.PathSeparator))
	return strings.Join(split, "/")
}
