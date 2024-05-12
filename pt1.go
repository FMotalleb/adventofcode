package main

import "strings"

var (
	vowels        [5]rune   = [5]rune{'a', 'e', 'i', 'o', 'u'}
	minimumVowels uint8     = 3
	badSequences  [4]string = [4]string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
)

func IsNicePt1(input string) bool {
	return hasVowels(input) && hasDuplicate(input) && !hasBadWords(input)
}

func hasVowels(input string) bool {
	counter := 0
	for _, i := range input {
		for _, o := range vowels {
			if i == o {
				counter++
				break
			}
		}
	}
	return counter >= 3
}

func hasDuplicate(input string) bool {
	for index, char := range input {
		if index == 0 {
			continue
		}
		if char == rune(input[index-1]) {
			return true
		}
	}
	return false
}
func hasBadWords(input string) bool {
	for _, i := range badSequences {
		if strings.Contains(input, i) {
			return true
		}
	}
	return false
}
