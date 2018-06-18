// Package numcn provides primitives for Interconversion
// between Chinese numbers & Numbers.
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

var cnSpecialUnits = []struct {
	cn  rune
	mul int64
}{
	// https://zh.wikipedia.org/wiki/%E4%B8%AD%E6%96%87%E6%95%B0%E5%AD%97
	// 在现代中文，“万进数”成为唯一的数字表示方式[j]，“上、中、下数”古法已不用，但仍有“兆应代表何值”的争议：
	// 兆的具体含义颇受争议，已知的可以表示1e6、1e12、1e16等，目前以1e12计算。
	{'京', 1e16},
	{'兆', 1e12},
	{'亿', 1e8},
	{'億', 1e8},
	{'万', 1e4},
	{'萬', 1e4},
}

var cnExtremeUnits = []struct {
	cn   rune
	unit float64
}{
	{'载', 1e44},
	{'正', 1e40},
	{'涧', 1e36},
	{'沟', 1e32},
	{'穰', 1e28},
	{'秭', 1e24},
	{'垓', 1e20},
	{'京', 1e16},
	{'兆', 1e12},
	{'亿', 1e8},
	{'億', 1e8},
	{'万', 1e4},
	{'萬', 1e4},
	{'又', 1e0},
	{'分', 1e-1},
	{'厘', 1e-2},
	{'毫', 1e-3},
	{'丝', 1e-4},
	{'忽', 1e-5},
	{'微', 1e-6},
	{'纤', 1e-7},
	{'沙', 1e-8},
	{'尘', 1e-9},
	{'纳', 1e-9},
	{'埃', 1e-10},
	{'渺', 1e-11},
	{'漠', 1e-12},
	{'皮', 1e-12},
	{'飞', 1e-15},
	{'阿', 1e-18},
	{'仄', 1e-21},
	{'幺', 1e-24},
}

const (
	cnNegativePrefix = rune('负')
	decimalPoint     = rune('点')
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
	for _, specialUnit := range cnSpecialUnits {
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
	0: '零',
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
			return string(append([]rune{cnNegativePrefix}, encodeWithoutYISHI(uint64(math.MaxInt64)+1)...))
		}
		return string(append([]rune{cnNegativePrefix}, encodeWithoutYISHI(uint64(-num))...))
	}
	return string(encodeWithoutYISHI(uint64(num)))
}

// encode Wrapper that remove "一十"
func encodeWithoutYISHI(num uint64) (res []rune) {
	ch := encodeFromInt64Helper(uint64(num))
	if len(ch) >= 2 && ch[0] == '一' && ch[1] == '十' {
		return ch[1:]
	}
	return ch
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
	return encodeSmallNum(num)
}

// encode number that less than 10000
func encodeSmallNum(num uint64) (res []rune) {
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

// MustDecodeToFloat64 : decode a chinese number string into Float64 without error
func MustDecodeToFloat64(cn string) float64 {
	res, err := DecodeToFloat64(cn)
	if err != nil {
		panic(err)
	}
	return res
}

// DecodeToFloat64 : decode a chinese number string into Float64
func DecodeToFloat64(cn string) (float64, error) {
	chars := []rune(cn)
	// positive or negative
	sign := float64(1)
	if chars[0] == cnNegativePrefix {
		sign = float64(-1)
		chars = chars[1:]
	}
	res, err := decodeToFloat64(chars)
	return res * sign, err
}

func decodeToFloat64(chars []rune) (res float64, err error) {
	// deal with decimal point
	for idx := len(chars) - 1; idx >= 0; idx-- {
		if chars[idx] == decimalPoint {
			integralPart, err := decodeToFloat64(chars[0:idx])
			if err != nil {
				return 0, err
			}
			decimalPart, err := decodeAfterDecimalPoint(chars[idx+1:])
			if err != nil {
				return 0, err
			}
			return integralPart + decimalPart, nil
		}
	}
	// deal with '万' & '亿'
	for _, specialUnit := range cnExtremeUnits {
		var left, right float64
		for idx := len(chars) - 1; idx >= 0; idx-- {
			if chars[idx] == specialUnit.cn {
				left, err = decodeToFloat64(chars[:idx])
				if err != nil {
					return 0, err
				}
				if idx == len(chars)-1 {
					right = 0
				} else {
					right, err = decodeToFloat64(chars[idx+1:])
					if err != nil {
						return 0, err
					}
				}
				return left*specialUnit.unit + right, nil
			}
		}
	}
	unit := float64(1) // current unit
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		if u, exist := cnUnit[char]; exist {
			unit = float64(u)
			if i == 0 {
				res += 10
			}
			continue
		}
		if n, exist := cnNum[char]; exist {
			res += unit * float64(n)
		} else {
			return 0, fmt.Errorf("confusing character %s", string(char))
		}
	}
	return res, nil
}

func decodeAfterDecimalPoint(chars []rune) (res float64, err error) {
	unit := float64(0.1)
	for _, char := range chars {
		if num, exist := cnNum[char]; exist {
			res += unit * float64(num)
		} else {
			return 0, fmt.Errorf("confusing character %s", string(char))
		}
		unit /= 10
	}
	return res, nil
}

// EncodeFromFloat64 : convert float64 into Chinese number
// 由于float64固然存在的精度问题，本函数可能不会特别精准。所以小数部分最多精确到6位。
func EncodeFromFloat64(num float64) string {
	if num < 0 {
		return string(cnNegativePrefix) + EncodeFromFloat64(-num)
	}
	integralPart := math.Floor(num)
	decimalPart := num - integralPart
	var integralRunes, decimalRunes, resRunes []rune
	if integralPart == 0 {
		integralRunes = []rune{'零'}
	} else {
		integralRunes = encodeFromFloat64Helper(integralPart)
	}
	if len(integralRunes) >= 2 && integralRunes[0] == '一' && integralRunes[1] == '十' {

		integralRunes = integralRunes[1:]
	}
	if decimalPart != 0 {
		decimalRunes = encodeDecimalPart(decimalPart)
		resRunes = append(integralRunes, decimalPoint)
		resRunes = append(resRunes, decimalRunes...)
	} else {
		resRunes = integralRunes
	}
	return string(resRunes)
}

// encodeHelper
func encodeFromFloat64Helper(num float64) (res []rune) {
	yi := num / 100000000
	if yi >= 1 {
		innerYi := num - math.Floor(num/100000000)*100000000
		left := append(encodeFromFloat64Helper(yi), '亿')
		right := encodeFromFloat64Helper(innerYi)
		// 若千位不为零，且万位为零；或者是千万位不为零，且亿位为零，则不需要补读零。
		// 例如：205,000读作“二十万五千”。
		if innerYi >= 1 && innerYi < 10000000 {
			left = append(left, '零')
		}
		return append(left, right...)
	}
	wan := num / 10000
	if wan >= 1 {
		innerWan := num - math.Floor(num/10000)*10000
		left := append(encodeFromFloat64Helper(wan), '万')
		right := encodeFromFloat64Helper(innerWan)
		// 若千位不为零，且万位为零；或者是千万位不为零，且亿位为零，则不需要补读零。
		// 例如：205,000读作“二十万五千”。
		if innerWan >= 1 && innerWan < 1000 {
			left = append(left, '零')
		}
		return append(left, right...)
	}
	return encodeSmallNum(uint64(num))
}

func encodeDecimalPart(num float64) []rune {
	num *= 1e7
	preciseNum := uint64(math.Round(num))
	res := make([]rune, 0)
	cnt := 6
	for preciseNum != 0 {
		if cnt == 0 {
			break
		}
		pos := uint64(1)
		for i := 0; i < cnt; i++ {
			pos *= 10
		}
		cur := preciseNum / pos
		res = append(res, numCn[uint64(cur)])
		preciseNum -= pos * cur
		cnt--
	}
	return res
}
