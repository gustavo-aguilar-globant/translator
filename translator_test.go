package translator_test

import (
	"path/filepath"
	"testing"

	"github.com/gustavo-aguilar-globant/translator"
)

func TestReadPropertiesFile(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		expectedResult map[string]string
	}{
		{
			name:     "ValidPropertiesFile",
			fileName: "es.properties",
			expectedResult: map[string]string{
				"hello":   "Hola",
				"goodbye": "Adiós",
				"welcome": "Bienvenido",
			},
		},
		{
			name:           "EmptyPropertiesFile",
			fileName:       "empty.properties",
			expectedResult: map[string]string{},
		},
		// Add more test cases if needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Read the properties file
			filePath := filepath.Join("testdata", test.fileName)
			properties, err := translator.ReadPropertiesFile(filePath)
			if err != nil {
				t.Fatalf("Failed to read properties file: %v", err)
			}

			// Verify the expected properties
			for key, expectedValue := range test.expectedResult {
				actualValue, ok := properties[key]
				if !ok {
					t.Errorf("Property key '%s' not found", key)
				} else if actualValue != expectedValue {
					t.Errorf("Unexpected value for property '%s'. Expected: %s, Actual: %s", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestTranslateText(t *testing.T) {
	translationProperties := map[string]string{
		"hello":   "Hola",
		"goodbye": "Adiós",
		"welcome": "Bienvenido",
	}

	tests := []struct {
		name                string
		sourceText          string
		expectedTranslation string
		expectedError       error
	}{
		{
			name:                "ValidTranslation",
			sourceText:          "goodbye",
			expectedTranslation: "Adiós",
			expectedError:       nil,
		},
		// Add more test cases for different source texts
	}

	for _, test := range tests {
		translation, err := translator.TranslateText(test.sourceText, translationProperties)

		// Compare the translation
		if translation != test.expectedTranslation {
			t.Errorf("Unexpected translation. Expected: %s, Actual: %s", test.expectedTranslation, translation)
		}

		// Compare the error
		if err != test.expectedError {
			t.Errorf("Unexpected error. Expected: %v, Actual: %v", test.expectedError, err)
		}
	}
}
