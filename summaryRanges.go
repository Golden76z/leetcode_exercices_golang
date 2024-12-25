package main

import "strconv"

func summaryRanges(nums []int) []string {
	finaltab := []string{}
	tempTab := []string{}
	var currentNum int
	if len(nums) < 1 {
		return finaltab
	} else {
		tempTab = append(tempTab, strconv.Itoa(nums[0]))
		currentNum = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != currentNum+1 {
			if len(tempTab) == 1 {
				finaltab = append(finaltab, tempTab[0])
			} else {
				finaltab = append(finaltab, tempTab[0]+"->"+tempTab[len(tempTab)-1])
			}
			tempTab = []string{}
			tempTab = append(tempTab, strconv.Itoa(nums[i]))
		} else {
			tempTab = append(tempTab, strconv.Itoa(nums[i]))
		}
		currentNum = nums[i]
	}
	if len(tempTab) == 1 {
		finaltab = append(finaltab, tempTab[0])
	} else {
		finaltab = append(finaltab, tempTab[0]+"->"+tempTab[len(tempTab)-1])
	}
	return finaltab
}
