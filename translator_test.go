package translator_test

import (
	"testing"

	"github.com/gustavo-aguilar-globant/translator"
)

func TestReadPropertiesFile(t *testing.T) {
	tests := []struct {
		filePath             string
		expectedTranslations map[string]string
		expectedError        error
	}{
		{
			filePath: "./testdata/es.properties",
			expectedTranslations: map[string]string{
				"hello":   "Hola",
				"goodbye": "Adiós",
				"welcome": "Bienvenido",
			},
			expectedError: nil,
		},
		{
			filePath: "./testdata/us.properties",
			expectedTranslations: map[string]string{
				"hello":   "Hello",
				"goodbye": "Goodbye",
				"welcome": "Welcome",
			},
			expectedError: nil,
		},
		// Add more test cases for different properties files
	}

	for _, test := range tests {
		translations, err := translator.ReadPropertiesFile(test.filePath)

		// Compare the error
		if err != test.expectedError {
			t.Errorf("Unexpected error. Expected: %v, Actual: %v", test.expectedError, err)
		}

		// Compare the translator
		if len(translations) != len(test.expectedTranslations) {
			t.Errorf("Unexpected number of translator. Expected: %d, Actual: %d", len(test.expectedTranslations), len(translations))
		}

		for key, expectedValue := range test.expectedTranslations {
			actualValue, ok := translations[key]
			if !ok {
				t.Errorf("Translation not found for key: %s", key)
			}

			if actualValue != expectedValue {
				t.Errorf("Unexpected value for key %s. Expected: %s, Actual: %s", key, expectedValue, actualValue)
			}
		}
	}
}

func TestTranslateText(t *testing.T) {
	translationProperties := map[string]string{
		"hello":   "Hola",
		"goodbye": "Adiós",
		"welcome": "Bienvenido",
	}

	tests := []struct {
		sourceText          string
		expectedTranslation string
		expectedError       error
	}{
		{
			sourceText:          "goodbye",
			expectedTranslation: "Adiós",
			expectedError:       nil,
		},
		{
			sourceText:          "welcome",
			expectedTranslation: "Bienvenido",
			expectedError:       nil,
		},
		// Add more test cases for different source texts
	}

	for _, test := range tests {
		translation, err := translator.TranslateText(test.sourceText, translationProperties)

		// Compare the error
		if err != test.expectedError {
			t.Errorf("Unexpected error. Expected: %v, Actual: %v", test.expectedError, err)
		}

		// Compare the translation
		if translation != test.expectedTranslation {
			t.Errorf("Unexpected translation. Expected: %s, Actual: %s", test.expectedTranslation, translation)
		}
	}
}
