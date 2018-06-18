[![MIT](https://img.shields.io/github/license/pkumza/numcn.svg)](https://github.com/pkumza/numcn/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/pkumza/numcn?status.svg)](https://godoc.org/github.com/pkumza/numcn)
[![Go Report Card](https://goreportcard.com/badge/github.com/pkumza/numcn)](https://goreportcard.com/report/github.com/pkumza/numcn)


# numcn

Interconversion between Chinese numbers &amp; Numbers. 
中文数字与数字的相互转换。

# Example

```go
func main() {
	chNum := "负十七亿零五十三万七千零一十六"
	num, _ := numcn.DecodeToInt64(chNum)
	fmt.Println(num) // -1700537016
	chNumAgain := numcn.EncodeFromInt64(num)
	fmt.Println(chNumAgain) // 负十七亿零五十三万七千零一十六

	chFloatNum := "负零点零七三零六"
	fNum, _ := numcn.DecodeToFloat64(chFloatNum)
	fmt.Printf("%f\n", fNum) // -0.073060
	chFloatNumAgain := numcn.EncodeFromFloat64(fNum)
	fmt.Println(chFloatNumAgain) // 负零点零七三零六
}
```
