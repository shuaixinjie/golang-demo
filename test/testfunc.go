package test

//测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头
//func TestSin(t *testing.T) { /* ... */ }
//其中t参数用于报告测试失败和附加的日志信息

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
