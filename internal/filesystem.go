package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func ParseFileContent(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("Error resolving path: %v", err)
	}

	file, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("File does not exist at %s\n", absPath)
	}
	if file.IsDir() {
		return "", fmt.Errorf("%s is a directory, not a file", absPath)
	}

	bytes, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}
	return string(bytes), nil
}
