package main

func longestCommonPrefix(strs []string) string {
	var result string
	cond := true
	var min int = 500
	if len(strs) < 1 {
		return result
	}
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return result
		} else if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	for i := 0; i < min; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				cond = false
			}
		}
		if cond {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
}
