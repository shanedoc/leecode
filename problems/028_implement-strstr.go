package problems

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
//输入：haystack = "hello", needle = "ll"
//输出：2

func StrStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 || l1 == 0 {
		return 0
	}
	j := 0
	for i := 0; i < l1; i++ {
		if haystack[i] == needle[j] {
			j++
			if j == l2 {
				return i - j + 1
			}
		} else {
			//i返回初始位置，j清空进行下次循环比较
			i = i - j
			j = 0
		}
	}
	return -1
}
