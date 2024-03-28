package project

type Project struct {
	Name string
}

func New(name string) *Project {
	return &Project{
		Name: name,
	}
}
