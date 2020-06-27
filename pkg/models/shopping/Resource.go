package shopping

import (
	"github.com/jinzhu/gorm"
)

type Resource struct {
	// 1=Banner图 2=商品
	Type       int    `gorm:"column:type;type:tinyint(4);not null;default:0"`
	Kind       int    `gorm:"column:kind;type:tinyint(4);not null;default:0"`
	ReferId    int    `gorm:"column:refer_id;type:int(10);not null;default:0"`
	Sort       int    `gorm:"column:sort;type:int(10);not null;default:0"`
	Url        string `gorm:"column:url;type:varchar(255);not null;default:''"`

	BaseModel
}

func (this *Resource) TableName() string {
	return "resources"
}

func (this *Resource) Condition(data map[string]interface{}, db *gorm.DB) {
	if value, ok := data["id"]; ok {
		db = db.Where("id = ?", value)
	}
}
