package numcn

import (
	"testing"
)

func TestNumCn(t *testing.T) {
	testCases := []struct {
		cn  string
		num int64
	}{
		{"零", 0},
		{"〇", 0},
		{"七", 7},
		{"七十八", 78},
		{"七百六十九", 769},
		{"柒仟壹佰肆拾", 7140},
		{"贰亿叁仟肆佰贰拾叁万肆仟贰佰贰拾叁", 234234223},
		{"七千零一", 7001},
		{"一万零一十四", 10014},
		{"十六万", 160000},
		{"十六", 16},
		{"六十", 60},
		{"负七百六十三", -763},
		{"一万四千六百二十三", 14623},
		{"一百零二亿五千零一万一千零三十八", 10250011038},
		{"一万一千一百一十一亿一千一百二十三万四千五百六十七", 1111111234567},
		{"五十万亿", 50000000000000},
		{"一亿亿", 10000000000000000},
		{"五十万亿零一百万零二十四", 50000001000024},
		{"负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八", -9223372036854775808},
		{"廿一", 21},
		{"卌玖", 49},
	}
	for _, testCase := range testCases {
		res, err := DecodeToInt64(testCase.cn)
		if err != nil {
			t.Fatalf("Error Happened %v", err)
		}
		if res != testCase.num {
			t.Errorf("ERROR!!! CN %s, Res %d, Num %d", testCase.cn, res, testCase.num)
		} else {
			t.Logf("CN %s, Res %d, Num %d", testCase.cn, res, testCase.num)
		}
	}
}

func TestEncode(t *testing.T) {
	testCases := []struct {
		cn  string
		num int64
	}{
		{"零", 0},
		{"一", 1},
		{"七十八", 78},
		{"七百六十九", 769},
		{"七千一百四十", 7140},
		{"七千零一", 7001},
		{"一万零一十四", 10014},
		{"十六", 16},
		{"六十", 60},
		{"负七百六十三", -763},
		{"一万四千六百二十三", 14623},
		{"一百零二亿五千零一万一千零三十八", 10250011038},
		{"一万一千一百一十一亿一千一百二十三万四千五百六十七", 1111111234567},
		{"五十万亿", 50000000000000},
		{"一亿亿", 10000000000000000},
		{"十六万", 160000},
		{"负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八", -9223372036854775808},
		{"五十万亿零一百万零二十四", 50000001000024},
	}
	for _, testCase := range testCases {
		res := EncodeFromInt64(testCase.num)
		if string(res) != testCase.cn {
			t.Errorf("ERROR!!! CN %s, Res %s, Num %d", testCase.cn, string(res), testCase.num)
		} else {
			t.Logf("CN %s, Res %s, Num %d", testCase.cn, string(res), testCase.num)
		}
	}
}
