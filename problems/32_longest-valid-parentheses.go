package problems

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//利用栈

func LongestValidParentheses(s string) int {
	ret, i, stack := 0, 0, []int{}
	//初始化，开始节点为)
	stack = append(stack, -1)
	if s == "" {
		return ret
	}
	for ; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			//弹出
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ret = max(ret, i-stack[len(stack)-1])
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
