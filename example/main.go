package main

import (
	"fmt"

	"github.com/pkumza/numcn"
)

func main() {
	chnum := "五十三万零一百七十六"
	num, _ := numcn.DecodeToInt64(chnum)
	fmt.Println(num)
}
