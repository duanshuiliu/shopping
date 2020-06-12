package models

import (
	"shopping/utils/orm"
)

type modeler interface{
	Table()
}

type BaseModel struct {
	// 
}

func (this *BaseModel) Create(model modeler, ) {
	// 
}

func (this *BaseModel) Update() {
	// 
}
