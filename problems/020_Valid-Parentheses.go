package problems

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。

func IsValid(s string) bool {
	n := len(s)
	if n == 1 || n%2 != 0 {
		return false
	}
	signal := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if signal[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != signal[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
