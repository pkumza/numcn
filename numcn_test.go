package numcn

import (
	"testing"
)

func TestDecodeToInt64(t *testing.T) {
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
		{"负十六", -16},
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

func TestEncodeFromInt64(t *testing.T) {
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
		{"负十六", -16},
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

func TestDecodeToFloat64(t *testing.T) {
	testCases := []struct {
		cn  string
		num float64
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
		{"一亿零一亿", 10000000100000000},
		{"五十万亿零一百万零二十四", 50000001000024},
		{"负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八", -9223372036854775808},
		{"廿一", 21},
		{"卌玖", 49},
		{"五十涧", 5e37},
		{"十又四厘九毫", 10.049},
		{"一万四千又四忽九纳", 14000.000040009},
		{"三点一四一五九二六", 3.1415926},
		{"壹仟陆佰零叁点伍玖陆", 1603.596},
	}
	for _, testCase := range testCases {
		res, err := DecodeToFloat64(testCase.cn)
		if err != nil {
			t.Fatalf("Critical Error Happened: %v", err)
		}
		if res != testCase.num {
			t.Errorf("Error: CN %s, Res %9f, Should be %9f", testCase.cn, res, testCase.num)
		} else {
			t.Logf("CN %s, Res %v, Num %v", testCase.cn, res, testCase.num)
		}
	}
}

func TestEncodeFromFloat64(t *testing.T) {
	testCases := []struct {
		cn  string
		num float64
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
		{"零点零零零零三七", 0.000037},
		{"三点一四一五九二", 3.141592},
		{"负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八", -9223372036854775808},
		{"五十万亿零一百万零二十四", 50000001000024},
		{"负十点一三五七九", -10.13579},
	}
	for _, testCase := range testCases {
		res := EncodeFromFloat64(testCase.num)
		if string(res) != testCase.cn {
			t.Errorf("ERROR!!! CN %s, Res %s, Num %f", testCase.cn, string(res), testCase.num)
		} else {
			t.Logf("CN %s, Res %s, Num %f", testCase.cn, string(res), testCase.num)
		}
	}
}
