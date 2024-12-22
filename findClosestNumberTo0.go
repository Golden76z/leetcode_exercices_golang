package main

import (
	"fmt"
	"math"
)

func main() {
	tab := []int{-4, -2, 1, -1, 4, 8}
	fmt.Println(FindClosestNumber(tab))
}

func FindClosestNumber(nums []int) int {
	var result int
	if len(nums) == 0 {
		return 0
	} else {
		result = nums[0]
	}
	for i := 0; i < len(nums); i++ {
		if math.Abs(float64(nums[i])) <= math.Abs(float64(result)) {
			if math.Abs(float64(nums[i])) == math.Abs(float64(result)) {
				if nums[i] > result {
					result = nums[i]
				}
			} else {
				result = nums[i]
			}
		}
	}
	return result
}
