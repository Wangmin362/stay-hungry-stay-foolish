package sync

import (
	"fmt"
	"os"
	"strings"
)

const (
	url          = "https://%s.%s/%s" // https://<bucketName>.<endpoint>/<path>
	syncImageDir = "vx_images"
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
