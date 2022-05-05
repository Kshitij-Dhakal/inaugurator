package common

import (
	"fmt"
	"strings"
)

func toDesiredCase(str string) string  {
	if(str == ""){
		return ""
	}
	var result string
	return result
}

func splitString(str string) []string   {
	if strings.Contains(str, "-") {
		return strings.Split(str, "-")
	}
	if strings.Contains(str, "_") {
		return strings.Split(str, "_")
	}
	if strings.Contains(str, ".") {
		return strings.Split(str, ".")
	}
	var result []string
	for _, v := range str {
		fmt.Println(v)
	}
	return result
}

func splitCamelCaseString(str string) []string {
	if(str == ""){
		return []string{}
	}
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