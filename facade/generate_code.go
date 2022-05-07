package facade

import (
	"errors"

	"github.com/Kshitij-Dhakal/inaugurator/service"
)

type CodeGenerator interface {
	GenerateCode(file string, args ...string) error
}

type CodeGeneratorImpl struct {
	Service service.CodeGenerator;
}

func (c CodeGeneratorImpl) GenerateCode(file string, args ...string) error {
	if len(args) != 2 {
		return errors.New("No name provided")
	}
	//type := args[0]
	name := args[1]
	return c.Service.GenerateCode(file, name)
}