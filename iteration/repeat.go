package iteration

import "strings"

func Repeat(character string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}

func Count(substr, str string) int {
	return strings.Count(str, substr)
}

func Index(substr, str string) int {
	return strings.Index(str, substr)
}

func Join(words []string, sep string) string {
	return strings.Join(words, sep)
}
