package parser

import (
	"encoding/json"
	"os"
	"strings"
)

// LoadCapitals reads the provincial capitals map from a JSON config file.
// Satisfies requirement: "provided at runtime via configuration."
func LoadCapitals(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var capitals map[string]string
	if err := json.Unmarshal(data, &capitals); err != nil {
		return nil, err
	}
	return capitals, nil
}

// EnrichCity appends province name if city is a provincial capital.
// Capitals map is injected as parameter — no global variable.
func EnrichCity(city string, capitals map[string]string) string {
	key := strings.ToUpper(strings.TrimSpace(city))
	if prov, ok := capitals[key]; ok {
		return key + " " + prov
	}
	return key
}
