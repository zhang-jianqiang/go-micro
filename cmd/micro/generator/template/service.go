package template

var ServiceSRV = `package service

import (
	"{{.Vendor}}{{.Service}}/domain/model"
	"{{.Vendor}}{{.Service}}/domain/repository"
)

type I{{title .Service}}DataService interface {
	FindOneById(int) (*model.{{title .Service}}, error)
}

func New{{title .Service}}DataService({{.Service}}Repository I{{title .Service}}Repository) {{title .Service}}DataService{
	return &{{title .Service}}DataService{ {{.Service}}Repository }
}

type {{title .Service}}DataService struct {
	{{title .Service}}Repository repository.I{{title .Service}}Repository
}

func (u *{{title .Service}}DataService) FindOneById(id int) (*model.{{title .Service}}, error) {
	return u.{{title .Service}}Repository.FindOneById(id)
}
`
