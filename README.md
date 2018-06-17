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
}
```
