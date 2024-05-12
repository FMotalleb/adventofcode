package main

func IsNicePt2(input string) bool {
	return hasRepeat(input) && hasLetterDiff(input)
}

func hasRepeat(input string) bool {
	breaker := len(input) - 1
	for i := range input {
		if i >= breaker {
			break
		}
		for j := range input {
			if j <= i+1 || j >= breaker {
				continue
			}
			if input[i] == input[j] && input[i+1] == input[j+1] {
				return true
			}

		}
	}
	return false
}

func hasLetterDiff(input string) bool {
	for index, char := range input {
		if index <= 1 {
			continue
		}
		if char == rune(input[index-2]) {
			return true
		}
	}
	return false
}
