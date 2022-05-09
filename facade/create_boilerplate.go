package facade

import "github.com/Kshitij-Dhakal/inaugurator/service"

type BoilerplateCreator interface {
	CreateBoilerplate(file string) (string, error)
}

type BoilerplateCreatorImpl struct {
	Service service.BoilderplaterCreator
}

func (c BoilerplateCreatorImpl) CreateBoilerplate(args ...string) (string, error) {
	//TODO: validation logic here
	return c.Service.CreateBoilerplate(args[0])
}
