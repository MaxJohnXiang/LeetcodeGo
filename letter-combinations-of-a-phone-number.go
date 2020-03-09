// package problem0017
//
// import (
// 	"strings"
// )

var m = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

//                2
//		a           b               c
//	ad ae af    bd be bf        cd ce cf

func letterCombinations(digits string) []string {
	//corner case
	if digits == "" {
		return nil
	}
	res := make([]string, 0)
	var dfs func(string, []string)

	dfs = func(temp string, digits []string) {
		if len(digits) == 0 {
			res = append(res, temp)
			return
		}
		for _, v := range m[digits[0]] {
			dfs(temp+v, digits[1:])
		}
	}
	dfs("", strings.Split(digits, ""))
	return res
}
