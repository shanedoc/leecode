package problems

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	len := len(arr)
	for i, j := 0, len-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
