package main

import (
	"github.com/gin-gonic/gin"
	"fmt"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

	route  "shopping/routes"
	config "shopping/utils/conf"
	orm    "shopping/utils/orm"
)

func init() {
	if err := config.Register(); err != nil {
		fmt.Println("配置文件加载失败", err)
	}

    //配置使用示例
	//app,err1      := config.New("app")
	//database,err2 := config.New("database")

	//if err1 != nil {
	//	fmt.Println("app conf:", err1)
	//}
	//if err2 != nil {
	//	fmt.Println("database conf:", err2)
	//}

	//fmt.Println(app.String("test01"), database.String("demo01::test01"), app.String("test02"))

	if err := orm.Register(); err != nil {
		fmt.Println("数据库初始化失败", err)
	}

	// 数据库使用示例
	//db, err1 := sql.Open("mysql", "root:*Kelvin2020@tcp(127.0.0.1:3306)/db01")

	//if err1 != nil {
	//	fmt.Println("数据库连接失败", err1)
	//}

	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(10)
	//err2 := db.Ping()

	//if err2 != nil {
	//	fmt.Println("DB ping error: ", err2)
	//}
}

func main() {
	r := gin.New()

	route.AddRoute(r)
	r.Run(":10003")
}
