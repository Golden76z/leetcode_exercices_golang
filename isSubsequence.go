package main

func isSubsequence(s string, t string) bool {
	count := 0
	if len(s) < 1 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[count] {
			count++
		}
		if count == len(s) {
			break
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
