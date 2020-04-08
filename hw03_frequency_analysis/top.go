package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

const count = 10

type KeyValue struct {
	Key   string
	Value int
}

func sortMapStringInt(m map[string]int) []KeyValue {
	result := make([]KeyValue, 0, len(m))
	for k, v := range m {
		result = append(result, KeyValue{k, v})
	}
	sort.Slice(result, func(i int, j int) bool {
		return result[i].Value > result[j].Value
	})
	return result
}

func Top10(s string) []string {
	regex := regexp.MustCompile(`(\p{L}(-)?)+`)
	result := make([]string, 0, count)
	words := regex.FindAllString(strings.ToLower(s), -1)
	wordsCount := make(map[string]int)

	for _, word := range words {
		wordsCount[word]++
	}

	rankedWords := sortMapStringInt(wordsCount)

	for i, w := range rankedWords {
		if i == count {
			break
		}
		result = append(result, w.Key)
	}

	return result
}
