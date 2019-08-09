package main

import (
	"strconv"
	"strings"
)

//func main() {
//	var input string
//	_, _ = fmt.Scanln(&input)
	//var result int
	//fmt.Println(calculate(input))
//}
func splitter(input string) []string {
	return strings.FieldsFunc(input, Split)
}
func Split(r rune) bool {
	return r == '+' || r == '-'
}
func Calculate(input string) int {
	elements := splitter(input)
	var index int8
	var result int
	result = 0
	for _, element := range elements {
		element = strings.TrimSpace(element)
		num,_ := strconv.ParseInt(element,10,64)
		index = int8(strings.Index(input, element))
		if index > 0 {
			switch input[index-1] {
			case '-':
				num *=-1
			case '+':
				//nothing :)
			}
		}
		result+=int(num)
	}
	return result
}
