package utils

import (
	"errors"
	"strings"
)

func SanitizeTitle(title string) (string, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return "", errors.New("title cannot be empty")
	}

	words := strings.Fields(title)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " "), nil
}