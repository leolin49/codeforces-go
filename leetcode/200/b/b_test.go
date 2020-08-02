// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[2,1,3,5,4,6,7]`, `2`, 
			`5`,
		},
		{
			`[3,2,1]`, `10`, 
			`3`,
		},
		{
			`[1,9,8,2,3,7,6,4,5]`, `7`, 
			`9`,
		},
		{
			`[1,11,22,33,44,55,66,77,88,99]`, `1000000000`, 
			`99`,
		},
		// TODO 测试参数的下界和上界
		{
			`[1,2]`,`1`,
			`2`,
		},
		{
			`[2,1]`,`1`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getWinner, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-200/problems/find-the-winner-of-an-array-game/
