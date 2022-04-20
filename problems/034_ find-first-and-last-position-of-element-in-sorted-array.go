package problems

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。

//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]

func SearchRange(nums []int, target int) []int {
	l := getLeftBorder(nums, target)
	if l == -1 {
		return []int{-1, -1}
	}
	r := getRightBorder(nums, target)
	if r-l > 1 {
		return []int{l + 1, r - 1}
	}
	return []int{-1, -1}
}

//分别查找左右边界的下标值
func getLeftBorder(nums []int, target int) int {
	border := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
			border = r
		} else {
			l = mid + 1
		}
	}
	return border
}

func getRightBorder(nums []int, target int) int {
	border := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
			border = l
		} else {
			r = mid - 1
		}
	}
	return border
}
