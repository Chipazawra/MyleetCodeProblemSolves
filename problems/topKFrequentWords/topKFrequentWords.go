//https://leetcode.com/problems/top-k-frequent-words/submissions/
package topkfrequentwords

import "sort"

func TopKFrequent(words []string, k int) []string {
	return topKFrequent(words, k)
}

func topKFrequent(words []string, k int) []string {

	ht := make(map[string]int)
	keys := make([]string, 0)
	for _, word := range words {
		ht[word]++

	}

	for key, _ := range ht {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return ht[keys[i]] > ht[keys[j]] || (ht[keys[i]] == ht[keys[j]] && keys[i] > keys[j])
	})

	return keys[:k]
}
