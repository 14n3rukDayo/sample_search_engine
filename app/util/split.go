package util

import "regexp"

func Split(input string) []string {
	re := regexp.MustCompile(`\s+`)
	words := re.Split(input, -1)

	var result []string
	for _, word := range words {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
