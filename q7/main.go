package main

import (
	"fmt"
	"log"
	"os"

	nb "github.com/takasho-tutrial/golang-check/q7/number"
)

func init() {
	nb.SetLogger(log.New(os.Stdout, "NumNum: ", log.Llongfile))
}

func main() {
	number, _ := nb.NewNumber(10)

	fmt.Printf("GetValue: %v\n", number.GetValue())
	fmt.Printf("Addition: %v\n", number.Addition(5))
	ans, _ := number.Division(3)
	fmt.Printf("Division: %v\n", ans)
	fmt.Printf("SetValue: %v\n", number.SetValue(100))
	fmt.Printf("GetValue: %v\n", number.GetValue())
}
