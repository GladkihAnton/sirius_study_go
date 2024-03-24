package my_string

import (
	"strings"
)

// n - kol-vo strings
// m - min dlina strings
// o(m^2 * n)

func LongestCommonPrefix(inputStrings []string) string {
	prefix := ""
	curPrefix := ""
	for _, curRune := range inputStrings[0] { // o(m)
		curChar := string(curRune)
		curPrefix = curPrefix + curChar

		for _, str := range inputStrings { // o(n)
			if !strings.HasPrefix(str, curPrefix) { // o(m)
				return prefix
			}
		}
		prefix = curPrefix
	}

	return prefix
}

func longestCommonPrefix(strs []string) string {
	current := strs[0]
	for _, str := range strs { // o(n)
		for i := 0; i < len(current); i++ { // o(m)
			if i == len(str) || current[i] != str[i] { // o(1)
				current = current[:i]
				break
			}
		}
	}
	return current
}
