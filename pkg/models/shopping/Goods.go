package shopping

import (
	//"github.com/jinzhu/gorm"
)

type Goods struct {
	BaseModel
	
	Type       int    `gorm:"column:type;type:tinyint(4);not null;default:0"`
	CategoryId int    `gorm:"column:category_id;type:int(10);not null;default:0"`
	Name       string `gorm:"column:name;type:varchar(50);not null;default:''"`
	Desc       string `gorm:"column:desc;type:varchar(255);not null;default:''"`
	Price      int    `gorm:"column:price;type:int(10);not null;default:0"`
	Sku        int    `gorm:"column:sku;type:int(10);not null;default:0"`
	Cover      string `gorm:"column:cover;type:varchar(255);not null;default:''"`
}

func (this *Goods) TableName() string {
	return "goods"
}
