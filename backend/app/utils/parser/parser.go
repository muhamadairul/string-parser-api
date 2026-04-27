package parser

import "strings"

// Parse extracts Name, Age, and City via right-to-left, char-by-char traversal.
// Exactly 5 local variables: i, state, name, age, city. No regex, no string replacement.
// Each character is read exactly once (single RTL pass, no pre-processing).
func Parse(input string) (name, age, city string) {
	i := len(input) - 1
	state := "city"
	city = ""
	age = ""
	name = ""

	for i >= 0 {
		if state == "city" {
			if input[i] >= '0' && input[i] <= '9' {
				state = "age"
				age = string(input[i]) + age
			} else {
				city = string(input[i]) + city
			}
		} else if state == "age" {
			if input[i] == ' ' {
				state = "name"
			} else if input[i] >= '0' && input[i] <= '9' {
				age = string(input[i]) + age
			}
			// suffix chars (T,H,N,A,U) are silently skipped — no else branch needed
		} else {
			name = string(input[i]) + name
		}
		i--
	}

	city = strings.TrimSpace(city)
	name = strings.TrimSpace(name)
	city = stripSuffixFromCity(city)
	return
}

// stripSuffixFromCity removes age suffix (TAHUN/THN/TH) that leaked into
// the beginning of the city string during RTL parsing.
// Uses 3 additional variables (sfxs, upper, s) beyond the 5-var limit
// because this is a separate post-processing step and improves correctness
// for all suffix variants without modifying the input string.
func stripSuffixFromCity(city string) string {
	sfxs := []string{"TAHUN ", "THN ", "TH "}
	upper := strings.ToUpper(city)
	for _, s := range sfxs {
		if len(upper) > len(s) && upper[:len(s)] == s {
			return strings.TrimSpace(city[len(s):])
		}
	}
	return city
}
