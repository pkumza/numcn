package numcn

import (
	"strings"
)

var cnNum = map[rune]int{
	'〇': 0,
	'一': 1,
	'二': 2,
	'三': 3,
	'四': 4,
	'五': 5,
	'六': 6,
	'七': 7,
	'八': 8,
	'九': 9,

	'零': 0,
	'壹': 1,
	'贰': 2,
	'叁': 3,
	'肆': 4,
	'伍': 5,
	'陆': 6,
	'柒': 7,
	'捌': 8,
	'玖': 9,

	'貮': 2,
	'两': 2,
}
var cnUnit = map[rune]int{
	'十': 10,
	'拾': 10,
	'百': 100,
	'佰': 100,
	'千': 1000,
	'仟': 1000,
}

var cnSpecialUnit = []struct {
	cn  string
	mul int
}{
	{"兆", 1000000000000},
	{"亿", 100000000},
	{"億", 100000000},
	{"万", 10000},
	{"萬", 10000},
}

func decode(cn string) (res int) {
	for _, specialUnit := range cnSpecialUnit {
		var left, right int
		lastIdx := strings.LastIndex(cn, specialUnit.cn)
		if lastIdx != -1 { // this unit does not exist
			left = decode(cn[:lastIdx])
			right = decode(cn[lastIdx+1:])
			return left*specialUnit.mul + right
		}
	}
	chars := []rune(cn)
	unit := 1 // current unit
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		if u, exist := cnUnit[char]; exist {
			unit = u
			if i == 0 {
				res += 10
			}
		}
		if n, exist := cnNum[char]; exist {
			res += unit * n
		}
	}
	return res
}
