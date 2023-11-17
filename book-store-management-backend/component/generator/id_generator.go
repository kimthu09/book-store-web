package generator

type IdGenerator interface {
	GenerateId() (string, error)
	IdProcess(id *string) (*string, error)
}
