package utils

import "unicode"

func HasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasNumber(s string) bool {
	for _, c := range s {
		if unicode.IsNumber(c) {
			return true
		}
	}
	return false
}
