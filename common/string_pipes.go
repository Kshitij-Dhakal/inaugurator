package common

import "strings"

func splitAndMapToTitle(str string) []string {
	split := splitString(str)
	var result []string
	for _, v := range split {
		result = append(result, MakeFirstLetterUpperCase(strings.ToLower(v)))
	}
	return result
}

func ToCamelCase(str string) string {
	result := splitAndMapToTitle(str)
	return MakeFirstLetterLowerCase(strings.Join(result, ""))
}

func ToPascalCase(str string) string {
	result := splitAndMapToTitle(str)
	return strings.Join(result, "")
}

func ToSnakeCase(str string) string {
	split := splitString(str)
	var result []string
	for _, v := range split {
		result = append(result, strings.ToLower(v))
	}
	return strings.Join(result, "_")
}