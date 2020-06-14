package shopping

import (
	"shopping/utils/orm"
)

type BaseModel struct {
	orm.Model

	Status int `gorm:"column:status;type:tinyint(4);not null;default:0"`
}

func (this *BaseModel) Connection() string {
	return "default"
}
