package template

var RepositorySRV = `package repository

import (
	"{{.Vendor}}{{.Service}}/domain/model"
	"gorm.io/gorm"
)
type I{{title .Service}}Repository interface{
	FindOneById(int) (*model.{{title .Service}}, error)
}

func New{{title .Service}}Repository(db *gorm.DB) I{{title .Service}}Repository  {
	return &{{title .Service}}Repository{mysqlDb:db}
}

type {{title .Service}}Repository struct {
	mysqlDb *gorm.DB
}

func (u *{{title .Service}}Repository) FindOneById(id int) ({{.Service}} *model.{{title .Service}}, err error) {
	err = u.mysqlDb.Where("id = ?", id).First({{.Service}}).Error
	return
}
`
