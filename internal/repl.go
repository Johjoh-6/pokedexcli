package internal

import "strings"

func CleanInput(text string) []string {
	split := strings.Split(text, " ")
	words := make([]string, 0, len(split))
	for _, word := range split {
		if word != "" {
			// Remove whitespace and lowercase
			words = append(words, strings.ToLower(strings.TrimSpace(word)))
		}
	}
	return words
}
