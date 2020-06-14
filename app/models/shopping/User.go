package shopping

import (
	//"github.com/jinzhu/gorm"
)

type User struct {
	BaseModel
	
	Type       int    `gorm:"column:type;type:tinyint(4);not null;default:3"`
	Openid     string `gorm:"column:openid;type:varchar(100);not null;default:'';unique"`
	Username   string `gorm:"column:username;type:varchar(100);not null;default:''"`
	Password   string `gorm:"column:password;type:varchar(500);not null;default:''"`
	Sex        int    `gorm:"column:sex;type:tinyint(4);not null;default:0"`
	Avatar     string `gorm:"column:avatar;type:varchar(255);not null;default:''"`
}

func (this *User) TableName() string {
	return "users"
}
