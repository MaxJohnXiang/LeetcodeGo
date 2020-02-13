// package problem0006
//
// import (
// 	"strings"
// )

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// dir
	flag := false

	curIndex := 0

	res := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		if curIndex == 0 || curIndex+1 == numRows {
			flag = !flag
		}

		res[curIndex] += string(s[i])

		if flag {
			curIndex++
		} else {
			curIndex--
		}

	}

	return strings.Join(res, "")
}
