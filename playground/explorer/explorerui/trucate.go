package explorerui

import "regexp"

func shortenWord(word string) string {
	if len(word) > 600 {
		return word[:10] + "..." + word[len(word)-10:]
	}
	return word
}

func processJSON(input string) string {
	wordRegex := regexp.MustCompile(`\S{601,}`)
	result := wordRegex.ReplaceAllStringFunc(input, shortenWord)
	return result
}
