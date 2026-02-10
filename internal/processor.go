package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func TransformContent(content string) (string, error) {
	dictionaryContent, err := os.ReadFile("./data/dictionary.yml")
	if err != nil {
		return "", fmt.Errorf("Error looking up dictionary: %v", err)
	}
	dictionary := make(map[string]string)

	err = yaml.Unmarshal(dictionaryContent, &dictionary)
	if err != nil {
		return "", fmt.Errorf("Error reading dictionary: %v", err)
	}

	re := regexp.MustCompile(`[a-zA-Z0-9']+`)
	result := re.ReplaceAllStringFunc(content, func(word string) string {
		if replacement, exists := dictionary[strings.ToUpper(word)]; exists {
			return replacement
		}
		return strings.ToUpper(word)
	})
	return result, nil
}
