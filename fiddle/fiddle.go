package main

import "fmt"

func main() {
	x := toSnakeCase("HelloWorld")
	fmt.Println(x)
	fmt.Println(len(x))
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

//make only first letter of string lowercase
func toLowerCase(str string) string {
	if(str == ""){
		return ""
	}
	var result string
	for _, v := range str {
		if v >= 65 && v <= 90 {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}