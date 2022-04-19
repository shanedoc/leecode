package problems

func Convert(s string, numRows int) string {
	//构造Z字形数组,遍历获取字符串
	ret, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			ret[down] = append(ret[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			ret[up] = append(ret[up], byte(s[i]))
			up--
			i++
		} else {
			down = 0
			up = numRows - 2
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range ret {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}
