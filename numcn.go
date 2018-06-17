package numcn

import (
	"fmt"
	"math"
)

var cnNum = map[rune]int64{
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
	'廿': 20,
	'卄': 20,
	'念': 20,
	'卅': 30,
	'卌': 40,
	'皕': 200,
}
var cnUnit = map[rune]int64{
	'十': 10,
	'拾': 10,
	'什': 10,
	'百': 100,
	'佰': 100,
	'陌': 100,
	'千': 1000,
	'仟': 1000,
	'阡': 1000,
}

var cnSpecialUnit = []struct {
	cn  rune
	mul int64
}{
	// https://zh.wikipedia.org/wiki/%E4%B8%AD%E6%96%87%E6%95%B0%E5%AD%97
	// 在现代中文，“万进数”成为唯一的数字表示方式[j]，“上、中、下数”古法已不用，但仍有“兆应代表何值”的争议：
	// 兆的具体含义颇受争议，已知的可以表示1e6、1e12、1e16等，目前以1e12计算。
	{'京', 10000000000000000},
	{'兆', 1000000000000},
	{'亿', 100000000},
	{'億', 100000000},
	{'万', 10000},
	{'萬', 10000},
}

const (
	cnNegativePrefix = rune('负')
)

// MustDecodeToInt64 : decode a chinese number string into Int64 without error
func MustDecodeToInt64(cn string) int64 {
	res, err := DecodeToInt64(cn)
	if err != nil {
		panic(err)
	}
	return res
}

// DecodeToInt64 : decode a chinese number string into Int64
func DecodeToInt64(cn string) (int64, error) {
	chars := []rune(cn)
	return decodeToInt64(chars)
}

func decodeToInt64(chars []rune) (res int64, err error) {
	// positive or negative
	sign := int64(1)
	if chars[0] == cnNegativePrefix {
		sign = int64(-1)
		chars = chars[1:]
	}
	// deal with '万' & '亿'
	for _, specialUnit := range cnSpecialUnit {
		var left, right int64
		for idx := len(chars) - 1; idx >= 0; idx-- {
			if chars[idx] == specialUnit.cn {
				left, err = decodeToInt64(chars[:idx])
				if err != nil {
					return 0, err
				}
				if idx == len(chars)-1 {
					right = 0
				} else {
					right, err = decodeToInt64(chars[idx+1:])
					if err != nil {
						return 0, err
					}
				}
				return sign * (left*specialUnit.mul + right), nil
			}
		}
	}
	unit := int64(1) // current unit
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		if u, exist := cnUnit[char]; exist {
			unit = u
			if i == 0 {
				res += 10
			}
			continue
		}
		if n, exist := cnNum[char]; exist {
			res += unit * n
		} else {
			return 0, fmt.Errorf("confusing character %s", string(char))
		}
	}
	return res * sign, nil
}

var numCn = map[uint64]rune{
	1: '一',
	2: '二',
	3: '三',
	4: '四',
	5: '五',
	6: '六',
	7: '七',
	8: '八',
	9: '九',
}

// EncodeFromInt64 : convert int64 into Chinese number
func EncodeFromInt64(num int64) string {
	if num == 0 {
		return "零"
	}
	if num < 0 {
		if num == math.MinInt64 {
			return string(append([]rune{cnNegativePrefix}, encodeFromInt64Helper(uint64(math.MaxInt64)+1)...))
		}
		return string(append([]rune{cnNegativePrefix}, encodeFromInt64Helper(uint64(-num))...))
	}

	ch := encodeFromInt64Helper(uint64(num))

	if len(ch) >= 2 && ch[0] == '一' && ch[1] == '十' {
		return string(ch[1:])
	}
	return string(ch)
}

// encode
func encodeFromInt64Helper(num uint64) (res []rune) {
	yi := num / 100000000
	if yi != 0 {
		innerYi := num % 100000000
		left := append(encodeFromInt64Helper(yi), '亿')
		right := encodeFromInt64Helper(innerYi)
		// 若千位不为零，且万位为零；或者是千万位不为零，且亿位为零，则不需要补读零。
		// 例如：205,000读作“二十万五千”。
		if innerYi != 0 && innerYi < 10000000 {
			left = append(left, '零')
		}
		return append(left, right...)
	}
	wan := num / 10000
	if wan != 0 {
		innerWan := num % 10000
		left := append(encodeFromInt64Helper(wan), '万')
		right := encodeFromInt64Helper(innerWan)
		// 若千位不为零，且万位为零；或者是千万位不为零，且亿位为零，则不需要补读零。
		// 例如：205,000读作“二十万五千”。
		if innerWan != 0 && innerWan < 1000 {
			left = append(left, '零')
		}
		return append(left, right...)
	}

	res = make([]rune, 0)
	qian := num / 1000
	bai := num / 100 % 10
	shi := num / 10 % 10
	ge := num % 10
	if qian != 0 {
		res = append(res, numCn[qian], '千')
	}
	if bai != 0 {
		res = append(res, numCn[bai], '百')
	} else {
		if qian != 0 && (shi != 0 || ge != 0) {
			res = append(res, '零')
		}
	}
	if shi != 0 {
		res = append(res, numCn[shi], '十')
	} else {
		if bai != 0 && ge != 0 {
			res = append(res, '零')
		}
	}
	if ge != 0 {
		res = append(res, numCn[ge])
	}
	return res
}
