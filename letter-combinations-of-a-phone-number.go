// package problem0017

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

func letterCombinations(digits string) []string {
	//corner acse
	if digits == "" {
		return nil
	}

	digs := strings.Split(digits, "")

	res := make([]string, 0)
	var dfs func(string, []string)
	dfs = func(temp string, digits []string) {
		if len(digits) == 0 {
			res = append(res, temp)
			return
		}
		charArr := m[digits[0]]
		digits = digits[1:]
		for _, v := range charArr {
			dfs(temp+v, digits)
		}
	}
	dfs("", digs)
	return res
}
