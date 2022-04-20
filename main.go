package main

import (
	_ "leetcode/problems"
)

func main() {
	s1 := "abcde"
	s2 := "ace"
	arr1 := make([]int32, len(s1))
	for i, v := range s1 {
		arr1[i] = v
	}

	m, n := len(s1)-1, len(s2)-1
	dp := [][]int{}
	res := make([]rune, dp[m][n])
	index := len(res) - 1
	for index >= 0 {
		if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else {
			res[index] = arr1[index]
			index--
			m--
			n--
		}
	}

}
