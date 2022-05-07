package common

import (
	"errors"
	"regexp"
	"strings"
)

func MakeFirstLetterLowerCase(str string) string {
	if str == "" {
		return ""
	}
	var result string
	for i := 0; i < len(str); i++ {
		v := str[i]
		if i == 0 {
			result += strings.ToLower(string(v))
		} else {
			result += string(v)
		}
	}
	return result
}

func MakeFirstLetterUpperCase(str string) string {
	if str == "" {
		return ""
	}
	var result string
	for i := 0; i < len(str); i++ {
		v := str[i]
		if i == 0 {
			result += strings.ToUpper(string(v))
		} else {
			result += string(v)
		}
	}
	return result
}

func ToDesiredCase(str string) (string, error) {
	if str == "" {
		return "", errors.New("String is empty")
	}
	r, _ := regexp.Compile("__[a-z]{2}__\\w+")
	if !r.MatchString(str) {
		return "", errors.New("Invalid string format. Proper string format : __[a-z]{2}__[a-zA-Z0-9]+")
	}
	split := Split(str, '_')
	dc := split[0]
	s := split[1]

	var result []string
	var joinWith string
	split = splitString(s)
	for _, v := range split {
		if dc == "sc" { //snake case
			result = append(result, strings.ToLower(v))
			joinWith = "_"
		} else if dc == "pc" { //pascal case
			result = append(result, MakeFirstLetterUpperCase(strings.ToLower(v)))
			joinWith = ""
		} else if dc == "cc" { //camel case
			result = append(result, MakeFirstLetterUpperCase(strings.ToLower(v)))
			joinWith = ""
		} else {
			return "", errors.New("Invalid desired case")
		}
	}
	finalString := strings.Join(result, joinWith)
	if dc == "cc" {
		finalString = MakeFirstLetterLowerCase(finalString)
	}
	return finalString, nil
}

func splitString(str string) []string {
	if strings.Contains(str, "-") {
		return strings.Split(str, "-")
	}
	if strings.Contains(str, "_") {
		return strings.Split(str, "_")
	}
	if strings.Contains(str, ".") {
		return strings.Split(str, ".")
	}
	if strings.Contains(str, " ") {
		return strings.Split(str, " ")
	}
	return splitCamelCaseString(str)
}

func splitCamelCaseString(str string) []string {
	if str == "" {
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

func Split(s string, r rune) []string {
	return strings.FieldsFunc(s, func(v rune) bool {
		return v == r
	})
}
