package main

import "fmt"

func productExceptSelf(nums []int) []int {
	finaltab := []int{}
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	var product_left, product_right, j int = 1, 1, len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(j)
		product_left *= nums[i]
		product_right *= nums[j]
		left[i+1] = product_left
		right[j-1] = product_right
		j--
	}
	fmt.Println(left)
	fmt.Println(right)
	for i := 0; i < len(nums); i++ {
		finaltab = append(finaltab, left[i]*right[i])
	}
	return finaltab
}

func productExceptSelf2(nums []int) []int {
	finaltab := []int{}
	var product int = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i {
				product *= nums[j]
			}
		}
		finaltab = append(finaltab, product)
		product = 1
	}
	return finaltab
}
