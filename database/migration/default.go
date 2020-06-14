package main

import (
	"fmt"
	"time"

	config    "shopping/utils/conf"
	orm       "shopping/utils/orm"
	mShopping "shopping/app/models/shopping"
)

var regModels []orm.ModelMaker

func init() {
	if err := config.Register(); err != nil {
		fmt.Println("配置文件加载失败", err)
	}

	if err := orm.Register(); err != nil {
		fmt.Println("数据库初始化失败", err)
	}
}

func main() {
	fmt.Println("开始创建数据库表单:") 

	// 注册数据库表单
	register(&mShopping.User{}, &mShopping.Category{})
	// 开始执行
	run()
}

func register(modelmaker ...orm.ModelMaker) {
	regModels = append(regModels, modelmaker...)
}

func run() {
	if regModels == nil {
		printMessage("暂无注册的模型")
		return
	}

	for _, model := range regModels {

		tableName := model.TableName()

		if err := createTable(model); err != nil {
			printMessage("表单创建失败", tableName, err)
		} else {
			printMessage("表单创建成功", tableName)
		}
	}
}

func createTable(modelmaker orm.ModelMaker) error {
	db, err := modelmaker.GetDb(modelmaker)

	if err != nil {
		return err
	} 

	result := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(modelmaker)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func printMessage(data ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), data)
}
