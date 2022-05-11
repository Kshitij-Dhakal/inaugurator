package boilerplate

type BoilderplateListerFacade interface {
	ListBoilerplates () error
}

type BoilerplateListerFacadeImpl struct {
	Service BoilderplateListerService
}

func (f BoilerplateListerFacadeImpl) ListBoilerplates() error {
	return f.Service.ListBoilerplates()
}