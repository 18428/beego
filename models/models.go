package models

import (
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id   int64
	Name string
	Age  int
}

func init() {
	orm.RegisterModel(new(Student))
}
