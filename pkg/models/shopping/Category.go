package shopping

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Category struct {
	Type       int8   `gorm:"column:type;type:tinyint(4);not null;default:0"`
	Pid        uint   `gorm:"column:pid;type:int(11);not null;default:0;"`
	Name       string `gorm:"column:name;type:varchar(50);not null;default:''"`
	Desc       string `gorm:"column:desc;type:varchar(255);not null;default:''"`

	BaseModel
}

func (this *Category) TableName() string {
	return "categories"
}

func (this *Category) Condition(data map[string]interface{}, db *gorm.DB) {
	if value, ok := data["id"]; ok {
		db = db.Where("id = ?", value)
	}
}
