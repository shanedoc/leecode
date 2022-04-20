package problems

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//输入：nums = [3,2,2,3], val = 3
//输出：2, nums = [2,2]
func RemoveElement(nums []int, val int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	//快慢指针：fast向后移动，如果fast！=val fast赋给slow
	//fast=val fast++
	j := 0
	for i := 0; i < len; i++ {
		if nums[i] != val {
			if i != j {
				//快慢指针交换数值
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}

	}
	return j
}
