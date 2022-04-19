package problems

//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
//[]string{"flower", "flow", "flight"}
func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]
	//依次比较数组中的字符串元素
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			//防止越界
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}
