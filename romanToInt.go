package main

func romanToInt(s string) int {
	var result int
	if len(s) == 1 {
		if s == "M" {
			return 1000
		} else if s == "D" {
			return 500
		} else if s == "C" {
			return 100
		} else if s == "L" {
			return 50
		} else if s == "X" {
			return 10
		} else if s == "V" {
			return 1
		} else if s == "I" {
			return 1
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 'C' {
			if s[i+1] == 'M' {
				result += 900
				i++
			} else if s[i+1] == 'D' {
				result += 400
				i++
			} else {
				result += 100
			}
		} else if s[i] == 'X' {
			if s[i+1] == 'C' {
				result += 90
				i++
			} else if s[i+1] == 'L' {
				result += 40
				i++
			} else {
				result += 10
			}
		} else if s[i] == 'I' {
			if s[i+1] == 'X' {
				result += 9
				i++
			} else if s[i+1] == 'V' {
				result += 4
				i++
			} else {
				result += 1
			}
		} else if s[i] == 'M' {
			result += 1000
		} else if s[i] == 'D' {
			result += 500
		} else if s[i] == 'L' {
			result += 50
		} else if s[i] == 'V' {
			result += 5
		}
		if i == len(s)-2 {
			if s[len(s)-1] == 'I' {
				result++
			} else if s[i+1] == 'V' {
				result += 5
			} else if s[i+1] == 'X' {
				result += 10
			} else if s[i+1] == 'L' {
				result += 50
			} else if s[i+1] == 'C' {
				result += 100
			} else if s[i+1] == 'D' {
				result += 500
			} else if s[i+1] == 'M' {
				result += 1000
			}
		}
	}
	return result
}
