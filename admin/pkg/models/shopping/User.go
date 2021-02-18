package shopping

import (
	"gorm.io/gorm"
)

type User struct {
	// 1=超级管理员 2=管理员 3=普通用户
	Type       int    `gorm:"column:type;type:tinyint(4);not null;default:3"`
	Phone      string `gorm:"column:phone;type:varchar(20);not null;default:'';unique"`
	Openid     string `gorm:"column:openid;type:varchar(100);not null;default:''"`
	Nickname   string `gorm:"column:nickname;type:varchar(50);not null;default:''"`
	Password   string `gorm:"column:password;type:varchar(255);not null;default:''"`
	Age        int    `gorm:"column:age;type:smallint(4);not null;default:0"`
	
	// 0=保密 1=男性 2=女性
	Sex        int    `gorm:"column:sex;type:tinyint(4);not null;default:0"`
	Avatar     string `gorm:"column:avatar;type:varchar(255);not null;default:''"`

	BaseModel
}

func (this *User) TableName() string {
	return "users"
}

func (this *User) Condition(data map[string]interface{}, db *gorm.DB) *gorm.DB {
	if value, ok := data["id"]; ok {
		db = db.Where("id = ?", value)
	}

	return db
}
