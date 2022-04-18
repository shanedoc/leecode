package problems

/*
给你一个字符串 s，找到 s 中最长的回文子串
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	//中心扩展法
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := getValue(s, i, i)
		l2, r2 := getValue(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func getValue(s string, l, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
