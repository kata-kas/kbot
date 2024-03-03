package normalization

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func Normalize(input string) string {
	// Normalize the input string to Unicode NFKD form
	normalized := norm.NFKD.String(input)

	result := polishToLatin(normalized)
	result = removeDiacritics(result)
	result = removeEmojis(result)

	return result
}

func NormalizeForComparison(input string) string {
	normalized := norm.NFKD.String(input)
	normalized = polishToLatin(normalized)

	var result []rune
	for _, char := range normalized {
		if unicode.IsLetter(char) && unicode.Is(unicode.Latin, char) && char != ' ' {
			result = append(result, unicode.ToLower(char))
		}
	}

	return string(result)
}

func polishToLatin(input string) string {
	// Replace Polish characters with Latin counterparts
	replacer := strings.NewReplacer(
		"ą", "a",
		"ć", "c",
		"ę", "e",
		"ł", "l",
		"ń", "n",
		"ó", "o",
		"ś", "s",
		"ź", "z",
		"ż", "z",
	)

	return replacer.Replace(input)
}

func removeDiacritics(input string) string {
	var result []rune
	for _, char := range input {
		if char == ' ' || (char <= 'z' && (char >= 'a' || char >= 'A')) {
			result = append(result, char)
		}
	}
	return string(result)
}

func removeEmojis(input string) string {
	emojiFilter := func(r rune) rune {
		if unicode.Is(unicode.So, r) || unicode.Is(unicode.Sk, r) {
			return -1 // Remove emojis and symbols
		}
		return r
	}

	result := strings.Map(emojiFilter, input)

	return result
}
