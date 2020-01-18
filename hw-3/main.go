package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(regexp.Match(`^\w$`, []byte(`,`)))
	fmt.Println(Top10("cat and ,dog one-two, - dog! two cats , and ?one man"))
	fmt.Println(Top10("one one one one one one one one one one two two two two three three three four five six seven eight nine ten eleven"))
}

type wordCount struct {
	word string
	count int
}

// Top10 функция подсчитывает количество повторений слов и возвращает их топ10 отсортированный по кол-ву повторов в порядке убывания.
func Top10(s string) []string {
	var result []string
	wordsCount := map[string]int{}
	words := strings.Split(s, " ")

	for _, word := range words {
		if !isPunctuation(word) {
			wordsCount[prepareWord(word)]++
		}
	}

	top := make([]wordCount, len(wordsCount))

	for word, count := range wordsCount {
		top = append(top, wordCount{
			word:  word,
			count: count,
		})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].count > top[j].count
	})

	topCount := 0

	if len(wordsCount) < 10 {
		topCount = len(wordsCount)
	} else {
		topCount = 10
	}

	for _, wordCount := range top[:topCount] {
		result = append(result, wordCount.word)
	}

	return result
}

func isPunctuation(s string) bool {
	res, _ := regexp.Match(`^\W$`, []byte(s))

	return res
}

func prepareWord(s string) string {
	// Очищаем от знаков препинания до и после слова.
	r, _ := regexp.Compile(`^[\W]*([\w-]*)[\W]*$`)
	return strings.ToLower(string(r.ReplaceAll([]byte(s), []byte(`$1`))))
}