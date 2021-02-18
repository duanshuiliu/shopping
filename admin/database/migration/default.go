package main

import (
	"fmt"
	"time"
	"os"

	config    "shopping/admin/pkg/conf"
	orm       "shopping/admin/pkg/orm"
	mShopping "shopping/admin/pkg/models/shopping"
)

var regModels []orm.ModelMaker

func init() {
	config.DefaultConfigPath = "../../conf"

	if err := orm.Register(); err != nil {
		fmt.Println("数据库初始化失败", err)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("开始创建数据库表单:")

	// 注册数据库表单
	regModels = append(regModels, &mShopping.User{}, &mShopping.Category{}, &mShopping.Goods{}, &mShopping.Resource{})
	// 开始执行
	run()
}

func run() {
	for _, model := range regModels {
		tableName := model.TableName()

		if err := createTable(model); err != nil {
			printMessage("表单创建失败", tableName, err)
		} else {
			printMessage("表单创建成功", tableName)
		}
	}
}

func createTable(m orm.ModelMaker) error {
	db, err := m.GetDb(m)
	if err != nil { return err }

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(m)
	if err != nil { return err }
	return nil
}

func printMessage(data ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), data)
}
