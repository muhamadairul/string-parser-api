package parser

import "strings"

// Parse extracts Name, Age, and City from string.
func Parse(input string) (name, age, city string) {
	input = normalizeSuffix(input)

	i := len(input) - 1
	state := "city"
	city = ""
	age = ""
	name = ""

	for i >= 0 {
		c := input[i]

		if state == "city" {
			if c >= '0' && c <= '9' {
				state = "age"
				age = string(c) + age
			} else {
				city = string(c) + city
			}
		} else if state == "age" {
			if c == ' ' {
				state = "name"
			} else if c >= '0' && c <= '9' {
				age = string(c) + age
			}
		} else {
			name = string(c) + name
		}

		i--
	}

	city = strings.TrimSpace(city)
	name = strings.TrimSpace(name)

	city = stripLeadingAgeSuffix(city)
	return
}

// stripLeadingAgeSuffix removes age suffix characters.
func stripLeadingAgeSuffix(city string) string {
	sfxs := []string{"TAHUN ", "THN ", "TH "}
	upper := strings.ToUpper(city)
	for _, s := range sfxs {
		if len(upper) > len(s) && upper[:len(s)] == s {
			return strings.TrimSpace(city[len(s):])
		}
	}
	return city
}

// normalizeSuffix normalizes suffix spacing.
func normalizeSuffix(s string) string {
	suffixes := []string{"TAHUN", "THN", "TH"}
	upper := strings.ToUpper(s)

	for _, sfx := range suffixes {
		sfxLen := len(sfx)
		sLen := len(upper)
		for j := 0; j < sLen-sfxLen; j++ {
			if upper[j:j+sfxLen] != sfx {
				continue
			}
			afterSfx := j + sfxLen
			if afterSfx < sLen && upper[afterSfx] != ' ' {
				continue
			}
			k := j - 1
			if k < 0 || upper[k] != ' ' {
				continue
			}
			k--
			if k < 0 || upper[k] < '0' || upper[k] > '9' {
				continue
			}
			s = s[:j-1] + s[j:]
			upper = upper[:j-1] + upper[j:]
			break
		}
	}
	return s
}
