package translator

import (
	"fmt"

	"github.com/magiconair/properties"
)

func ReadPropertiesFile(filePath string) (map[string]string, error) {
	p, err := properties.LoadFile(filePath, properties.UTF8)
	if err != nil {
		return nil, err
	}

	propertiesMap := make(map[string]string)
	for _, key := range p.Keys() {
		propertiesMap[key] = p.GetString(key, "")
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
