package facade

import "github.com/Kshitij-Dhakal/inaugurator/service"

type BoilerplateCreator interface {
	CreateBoilerplate(file string) (string, error)
}

type BoilerplateCreatorImpl struct {
	createBoilderplateService service.CreateBoilerplateService;
}

func (c BoilerplateCreatorImpl) CreateBoilerplate(file string) (string, error)  {
	//TODO: validation logic here
	return c.createBoilderplateService.CreateBoilerplate(file)
}