package main

import (
	"fmt"

	cnb "github.com/takasho-tutrial/golang-check/q7/complexNumber"
)

func main() {
	cNum, _ := cnb.NewComplexNumber(1, 2)
	fmt.Printf("%v\n", cNum.GetNorm())
}
