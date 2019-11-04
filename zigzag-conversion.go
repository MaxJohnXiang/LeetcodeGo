// package problem0006
//
// import (
// 	"strings"
// )

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	flag, cur := false, 0
	strArr := make([]string, numRows)
	for _, v := range s {
		strArr[cur] += string(v)
		if cur == 0 || cur == numRows-1 {
			flag = !flag
		}
		if flag {
			cur++
		} else {
			cur--
		}
	}
	return strings.Join(strArr, "")
}
