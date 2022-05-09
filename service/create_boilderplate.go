package service

type BoilderplaterCreator interface {
	CreateBoilerplate(file string) (string, error)
}


type BoilderplaterCreatorImpl struct {
	
}

func (*BoilderplaterCreatorImpl) CreateBoilerplate(file string) (string, error) {
	return "", nil
}