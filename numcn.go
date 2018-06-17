package numcn

import "fmt"

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
				return left*specialUnit.mul + right, nil
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
