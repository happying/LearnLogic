package logic

import "fmt"

func TestIsSubsequence() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence_v2(s, t)
	fmt.Println(result)
}

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

示例 1：

输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：

输入：s = "axc", t = "ahbgdc"
输出：false

提示：
0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。
*/

func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}
	if len(t) <= 0 {
		return false
	}
	sIndex := 0
	sCurrent := string(s[0])
	for _, tchar := range t {
		if sCurrent == string(tchar) {
			sIndex++
			if sIndex >= len(s) {
				return true
			}
			sCurrent = string(s[sIndex])
		}
	}
	return false
}

/*
方法二：动态规划
思路及算法
考虑前面的双指针的做法，我们注意到我们有大量的时间用于在 t 中找到下一个匹配字符。
这样我们可以预处理出对于 t 的每一个位置，从该位置开始往后每一个字符第一次出现的位置。
我们可以使用动态规划的方法实现预处理，令 f[i][j]表示字符串 t 中从位置 i 开始往后字符 j 第一次出现的位置。在进行状态转移时，如果 t 中位置 i 的字符就是 j，那么 f[i][j]=i ，否则 j 出现在位置 i+1 开始往后，即 f[i][j]=f[i+1][j]，因此我们要倒过来进行动态规划，从后往前枚举 i。

这样我们可以写出状态转移方程：

f[i][j]={i,t[i]=jf[i+1][j],t[i]≠jf[i][j]=\begin{cases} i, & t[i]=j\\ f[i+1][j], & t[i] \neq j \end{cases}
f[i][j]={
i,
f[i+1][j],
​

t[i]=j
t[i]

=j
​

假定下标从 0 开始，那么 f[i][j] 中有 0≤i≤m−10 \leq i \leq m-10≤i≤m−1 ，对于边界状态 f[m−1][..]，我们置 f[m][..] 为 m，让 f[m−1][..] 正常进行转移。这样如果 f[i][j]=m，则表示从位置 i 开始往后不存在字符 j。

这样，我们可以利用 f 数组，每次 O(1) 地跳转到下一个位置，直到位置变为 m 或 s 中的每一个字符都匹配成功。

同时我们注意到，该解法中对 t 的处理与 s 无关，且预处理完成后，可以利用预处理数组的信息，线性地算出任意一个字符串 s 是否为 t 的子串。这样我们就可以解决「后续挑战」啦。
*/

func isSubsequence_v2(s string, t string) bool {
	sLength, tLength := len(s), len(t)
	f := make([][26]int, tLength+1)
	for i := 0; i < 26; i++ {
		f[tLength][i] = tLength
	}
	for i := tLength - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			byteA := byte(j + 'a')
			str := string(byteA)
			fmt.Println(str)
			if t[i] == byteA {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	add := 0
	for i := 0; i < sLength; i++ {
		if f[add][int(s[i]-'a')] == tLength {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}
	return true
}
