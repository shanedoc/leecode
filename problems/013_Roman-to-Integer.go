package problems

//给定一个罗马数字，将其转换成整数
var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	ret := 0
	//通常情况下小的数字在大的数字的右边
	len := len(s)
	for i := 0; i < len; i++ {
		value := romanMap[s[i]]
		if i+1 < len && value < romanMap[s[i+1]] {
			ret -= value
		} else {
			ret += value
		}
	}
	return ret
}
