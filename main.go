package main

import (
	"fmt"
	"leetcode/problems"
)

func main() {
	p := []int{1, 0, -1, 0, -2, 2}
	r := problems.FourSum(p, 0)
	fmt.Println(r)
}
