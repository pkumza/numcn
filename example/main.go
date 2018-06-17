package main

import (
	"fmt"

	"github.com/pkumza/numcn"
)

func main() {
	chNum := "负十七亿零五十三万七千零一十六"
	num, _ := numcn.DecodeToInt64(chNum)
	fmt.Println(num) // -1700537016
	chNumAgain := numcn.EncodeFromInt64(num)
	fmt.Println(chNumAgain) // 负十七亿零五十三万七千零一十六
}
