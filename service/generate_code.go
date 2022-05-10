package service

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Kshitij-Dhakal/inaugurator/common"
)

type CodeGenerator interface {
	GenerateCode(file string, name string) error
}

type CodeGeneratorImpl struct {
}

func (c CodeGeneratorImpl) GenerateCode(file string, name string) error {
	split := common.Split(file, '/')
	fileNameSplit := common.Split(split[len(split)-1], '.')
	fileName := strings.Replace(fileNameSplit[0], "name", name, -1)

	newFileName, err := common.ToDesiredCase(fileName)
	if len(fileNameSplit) > 1 {
		fileNameSplit[0] = newFileName
		newFileName = strings.Join(fileNameSplit, ".")
	}
	split[len(split)-1] = newFileName
	if err != nil {
		return err
	}
	funcMap := map[string]interface{}{
		"ToPascalCase": common.ToPascalCase,
		"ToCamelCase":  common.ToCamelCase,
	}
	data, err := common.ReadFile(file)
	if err != nil {
		return err
	}

	template, err := template.New("code-template").Funcs(funcMap).Parse(data)
	if err != nil {
		return err
	}
	var processed bytes.Buffer
	err = template.Execute(&processed, name)
	fullPath := strings.Join(split, "/")
	fmt.Println("Generating file :", fullPath)
	return common.StoreString(fullPath, processed.String())
}
