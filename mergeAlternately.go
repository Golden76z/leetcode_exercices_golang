package main

func mergeAlternately(word1 string, word2 string) string {
	var result string
	longest := 0
	if len(word1) > len(word2) {
		longest = len(word1)
	} else {
		longest = len(word2)
	}
	for i := 0; i < longest; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}
	}
	return result
}
