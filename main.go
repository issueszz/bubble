package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)
func main() {
	// 检测
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./bubble conf/config.ini")
		return
	}

	// 读取配置信息
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed:%v", err)
		return
	}

	// 连接数据库
	err := dao.InitMySql(setting.Conf.MysqlConfig)

	if err != nil {
		fmt.Printf("init mysql failed: %v", err)
		return
	}

	dao.DB.AutoMigrate(&models.ToDo{})

	// 设置路由
	r := routers.SetupRouter()
	// 启动服务
	if err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		panic(err)
	}
}
