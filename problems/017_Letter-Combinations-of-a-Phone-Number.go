package problems

//定义全局变量
var (
	number = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	ret = []string{}
)

func LetterCombinations(digits string) []string {
	if len(digits) == 0 || digits == "" {
		return nil
	}
	//下标从0开始
	backtracking(digits, 0, "")
	return ret
}

func backtracking(digits string, index int, s string) {
	if index == len(digits) {
		ret = append(ret, s)
		return
	}
	num := digits[index] - '0'
	letter := number[num]
	for i := 0; i < len(letter); i++ {
		s = s + string(letter[i])
		backtracking(digits, index+1, s)
		s = s[:len(s)-1]
	}
}
