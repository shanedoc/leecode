package problems

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//请必须使用时间复杂度为 O(log n) 的算法

func SearchInsert(nums []int, target int) int {
	i := 0
	len := len(nums)
	//1、数组第一个元素大于target
	if nums[0] > target {
		return i
	}
	if nums[len-1] < target {
		return len
	}
	for ; i < len; i++ {
		//fmt.Println(i, nums[i])
		if nums[i] == target || nums[i] > target {
			break
		}
	}
	return i
}
