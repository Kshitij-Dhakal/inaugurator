package service

type CreateBoilerplateService interface {
	CreateBoilerplate(file string) (string, error)
}