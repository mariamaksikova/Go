package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}
	runes := []rune(word)
	first := runes[0]
	rest := runes[1:]
	for i, j := 0, len(rest)-1; i < j; i, j = i+1, j-1 {
		rest[i], rest[j] = rest[j], rest[i]
	}
	return string(first) + string(rest)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	encryptedWords := make([]string, len(words))
	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}
	return strings.Join(encryptedWords, " ")
}

func main() {
	phrases := []string{
		"Pepe Schnele is a legend",
		"Hello world",
		"Go programming is fun",
		"привет мир",
	}
	for _, p := range phrases {
		encrypted := encryptPhrase(p)
		fmt.Printf("Исходная: %s\nЗашифрованная: %s\n\n", p, encrypted)
	}
}
