// @Author Zihao_Li 2023/2/3 13:34:00
package mysql

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var Db *gorm.DB

func Init(cfg *settings.MySQLConfig) error {
	//grom 2。0 之后的连接方式
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil // panic(fmt.Errorf("gorm open error: %s\n", err))
	}
	//TODO:将表配置到数据库中去
	if err := db.AutoMigrate(
		&Video{},
	); err != nil {
		fmt.Println(err)
		return nil
	}
	a, err := db.DB()
	a.SetMaxOpenConns(cfg.MaxOpenConns)
	a.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil

}

func Close() {
	a, err := db.DB()
	err = a.Close()
	if err != nil {
		fmt.Println(err)
	}
}
