package main

import (
	"sort"
)

func GetSliceValue(nums []int) []int {
	//排序
	sort.Ints(nums)
	i, j := 0, 1
	//快慢指针：不相等把fast赋值给slow 相等fast++
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return nums[:i+1]
}
