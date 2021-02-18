package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"

	"shopping/admin/routes"
	"shopping/admin/pkg/conf"
	"shopping/admin/pkg/orm"
)

func init() {
	if err := orm.Register(); err != nil {
		fmt.Println("数据库初始化失败", err)
		os.Exit(0)
	}
}

func main() {
	appConf, err := conf.New("app")
	if err != nil {
		fmt.Println("获取配置文件失败", err)
		return
	}

	dsn := appConf.GetString("app.host")+":"+appConf.GetString("app.port")
	fmt.Println(dsn)

	r := gin.New()
	routes.AddRoute(r)
	r.Run(dsn)
}
