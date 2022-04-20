package problems

//搜索旋转排序数组
//时间复杂度O(log n)

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	//fmt.Println(nums)
	for l <= r {
		mid := l + (r-l)/2
		//fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}
		//判断target在那个哪个区间，每次舍弃另一半区间
		//判断端点值找到有序区间，判断target是否在有序区间中
		if nums[mid] >= nums[l] {
			if nums[l] <= target && target > nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		//fmt.Println(l, r)
	}
	return -1
}
