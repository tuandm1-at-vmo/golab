package lab

import "strings"

func CountWords(text string) map[string]int {
	words := strings.FieldsFunc(text, func(c rune) bool {
		return !strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", c)
	})
	presence := make(map[string]int)
	for _, word := range words {
		count := presence[word]
		presence[word] = count + 1
	}
	return presence
}
