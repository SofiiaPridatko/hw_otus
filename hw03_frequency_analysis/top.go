package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type entry struct {
	word  string
	count int
}

func Top10(str string) []string {
	words := strings.Fields(str)
	// Считаем количество слов
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	// Перекладываем в слайсы, чтобы менять последовательность
	entries := make([]entry, 0, len(freq))
	for w, c := range freq {
		entries = append(entries, entry{w, c})
	}
	// Сортировка по количеству и лексикографически
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count == entries[j].count {
			return entries[i].word < entries[j].word
		}
		return entries[i].count > entries[j].count
	})
	// Вывод только 10 элементов
	limit := 10
	if len(entries) < limit {
		limit = len(entries)
	}

	result := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		result = append(result, entries[i].word)
	}

	return result
}
