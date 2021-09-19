package dao

import (
	"bubble/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
)


func InitMySql(cfg *setting.MysqlConfig) (err error) {
	// dsn := "root:root123@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}