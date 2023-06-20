package translator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func ReadPropertiesFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open properties file: %w", err)
	}
	defer file.Close()

	propertiesMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	decoder := charmap.ISO8859_1.NewDecoder()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue // Skip comments and empty lines
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format in properties file: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		decodedValue, err := decoder.String(value)
		if err != nil {
			return nil, fmt.Errorf("failed to decode property value: %w", err)
		}

		propertiesMap[key] = decodedValue
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read properties file: %w", err)
	}

	return propertiesMap, nil
}

func TranslateText(sourceText string, translationProperties map[string]string) (string, error) {
	translatedText, ok := translationProperties[sourceText]
	if !ok {
		return "", fmt.Errorf("translation not found for text: %s", sourceText)
	}

	return translatedText, nil
}
