package gouroutine

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {
	counted := make(chan pair)
	stats := make(counter)

	i := 0
	go func() {
		for {
			word := next()
			i += 1
			if word == "" {
				counted <- pair{}
				break
			}
			counted <- pair{word, countDigits(word)}

		}
	}()

	for {
		countedPair := <-counted
		x := pair{}
		if countedPair == x {
			break
		}
		stats[countedPair.word] = countedPair.count
	}

	return stats
}

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func mainTask2() {
	phrase := "0ne 1wo thr33 4068 12312"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
