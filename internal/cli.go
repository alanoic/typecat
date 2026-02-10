package internal

import (
	"fmt"
	"os"
	"strings"
)

func ParseArgs(args []string) (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("Usage: typecat <path_to_file>")
	}

	var path string = strings.TrimSpace(os.Args[1])
	if path == "" {
		return "", fmt.Errorf("File path cannot be empty")
	}
	return path, nil
}
