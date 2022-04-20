package problems

//整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
//整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
//更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
//那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
//如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

func NextPermutation(nums []int) {
	//要使数字变大，任意一位变大即可：从后向前遍历从十位开始，十位右侧是否有大于十位的数字，没有则向左移动
	//从右向左找第一个不再递增的数字，从左向右找第一个好好大于当前位的数字
	len := len(nums)
	i := len - 2
	//从右向左找到第一个非递增数字
	for i >= 0 && nums[i+1] < nums[i] {
		i--
	}
	if i <= 0 {
		//直接倒序输出
		reverse(&nums, 0)
	}
	//找到刚好大于nums[i]的数字
	j := len - 1
	for j >= 0 && nums[j] >= nums[i] {
		j--
	}
	//交换i、j位置
	swap(&nums, i, j)
	//reverse(&nums, i+1)
}

func swap(nums *[]int, i int, j int) {
	tmp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = tmp
}

func reverse(nums *[]int, i int) {
	j := len(*nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
