package problems

//回溯
//注意左右括号出现的规则：先有左再有右
func GenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(l, r int, str string, res *[]string) {
	if l == 0 && r == 0 {
		//终止条件
		*res = append(*res, str)
		return
	}
	if l > 0 {
		findGenerateParenthesis(l-1, r, str+"(", res)
	}
	if r > 0 && r > l {
		findGenerateParenthesis(l, r-1, str+")", res)
	}
}
