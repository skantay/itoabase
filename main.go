package main

import "fmt"

func main() {
	value := -42
	base := 16
	result := ItoaBase(value, base)
	fmt.Printf("Integer %d in base %d: %s\n", value, base, result)
}

func ItoaBase(value, base int) string {
	var result string
	baseChars := getBase(base)

	for value != 0 {
		digitIndex := value % len(baseChars)
		value /= len(baseChars)
		char, err := getValueBase(digitIndex, baseChars)
		if !err {
			return ""
		}

		result = string(char) + result
	}

	return result
}

func getValueBase(digitIndex int, base string) (rune, bool) {
	for i, v := range base {
		if digitIndex == i {
			return v, true
		}
	}
	return '.', false
}

func getBase(base int) string {
	var result string
	var char rune = '0'
	var charL rune = 'A'

	for i := 0; i < base; i++ {
		if i > 9 {
			result += string(charL)
			charL++
			continue
		}
		result += string(char)
		char++
	}

	return result
}
