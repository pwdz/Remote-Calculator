package main

import (
	"fmt"
	"strings"
)
func main()  {
	var input string
	_, _ = fmt.Scanln(&input)
	//var result int
	//:= calculate()
}
func splitter(input string) []string  {
	return strings.FieldsFunc(input,Split)
}
func Split(r rune) bool {
	return r == '+' || r == '-'
}
func calculate() int {

	return 0
}
