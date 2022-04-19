package problems

import (
	"math"
	"sort"
)

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近
//返回这三个数的和。
//假定每组输入只存在恰好一个解。

func ThreeSumClosest(nums []int, target int) int {
	//双指针：从两侧向中间逼近
	//前后夹逼
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		//对数组进行排序
		sort.Ints(nums)
		//i从0顺序遍历
		//j=i+1,k=len(num)-1 像中间遍历
		for i := 0; i < n; i++ {
			//过滤重复元素
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum < target {
					j++
				} else {
					k--
				}
			}
		}

	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
