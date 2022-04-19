package problems

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。

func MaxArea(height []int) int {
	//双指针
	max, l, r := 0, 0, len(height)-1
	for l < r {
		wid := r - l
		high := 0
		//确定矩形高度
		if height[l] < height[r] {
			high = height[l]
			l++
		} else {
			high = height[r]
			r--
		}
		//fmt.Println(high, wid)
		tmp := high * wid
		if max < tmp {
			max = tmp
		}
	}
	return max
}
