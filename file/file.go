package file

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func IsJsonExtension(path string) bool {
	return strings.HasSuffix(path, ".json")
}
