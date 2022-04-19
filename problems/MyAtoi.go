package problems

import "math"

//请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）
/*
	函数 myAtoi(string s) 的算法如下：

	读入字符串并丢弃无用的前导空格
	检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
	读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
	将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
	如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
	返回整数作为最终结果。
*/

func MyAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//去掉空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = abs*10 + int(s[i]-'0') //字节byte '0'=46
		if abs*sign < math.MinInt32 {
			return math.MinInt32
		} else if abs*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return abs * sign
}
