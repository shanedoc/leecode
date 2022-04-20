package problems

//给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致
//nums = [1,1,2]
//2, nums = [1,2,_]

func RemoveDuplicates(nums []int) int {
	//快慢指针从0、1下标开始
	//i==j j++
	//i!=j nums[i+1]=nums[j] j++
	len := len(nums)
	if len <= 1 {
		return len
	}
	i := 0
	for j := 1; j < len; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
