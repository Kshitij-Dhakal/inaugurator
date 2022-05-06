package facade

type CodeGenerator interface {
	GenerateCode(id string, args ...string) (string, error)
}