package main

import (
	"fmt"
	"github.com/Kshitij-Dhakal/inaugurator/common"
)

func main() {
	fmt.Println(common.ToDesiredCase("__pc__helloWorld"))
	fmt.Println(common.ToDesiredCase("__cc__helloWorld"))
	fmt.Println(common.ToDesiredCase("__sc__helloWorld"))
	fmt.Println(common.ToDesiredCase("__helloWorld"))
	fmt.Println(common.ToDesiredCase("pc__helloWorld"))
	fmt.Println(common.ToDesiredCase("__pchelloWorld"))
	fmt.Println(common.ToDesiredCase("____helloWorld"))
}

//function to convert camelCase to snake_case
func toSnakeCase(str string) []string {
	var result []string
	var temp string
	for _, v := range str {
		if v >= 65 && v <= 90 {
			result = append(result, temp)
			temp = string(v)
		} else {
			temp += string(v)
		}
	}
	result = append(result, temp)
	return result
}

