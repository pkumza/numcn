package numcn

import (
	"testing"
)

func TestNumCn(t *testing.T) {
	testCases := []struct {
		cn  string
		num int
	}{
		{"零", 0},
		{"〇", 0},
		{"七", 7},
		{"七十八", 78},
		{"七百六十九", 769},
		{"柒仟壹佰肆拾", 7140},
		{"贰亿叁仟肆佰贰拾叁万肆仟贰佰贰拾叁", 234234223},
		{"七千零一", 7001},
		{"一百一十四", 114},
		{"十六", 16},
		{"六十", 60},
		{"一万四千六百二十三", 14623},
		{"一百零二亿五千零一万一千零三十八", 10250011038},
		{"一兆一千一百一十一亿一千一百二十三万四千五百六十七", 1111111234567},
		{"五十万亿", 50000000000000},
		{"一亿亿", 10000000000000000},
		{"一万兆", 10000000000000000},
		{"五十万亿零三千一百万零二十四", 50000031000024},
		{"三十万零七十兆零四百三十亿零四百零一万", 300070043004010000},
	}
	for _, testCase := range testCases {
		res := decode(testCase.cn)
		if res != testCase.num {
			t.Errorf("ERROR!!! CN %s, Res %d, Num %d", testCase.cn, res, testCase.num)
		} else {
			t.Logf("CN %s, Res %d, Num %d", testCase.cn, res, testCase.num)
		}
	}
}
