// @Author Zihao_Li 2023/2/3 13:34:00
package mysql

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	//TODO:将表配置到数据库中去
	if err := Db.AutoMigrate(
		&Video{}, &User{},
	); err != nil {
		fmt.Println(err)
		return nil
	}
	a, err := Db.DB()
	a.SetMaxOpenConns(cfg.MaxOpenConns)
	a.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil

}

func Close() {
	a, err := Db.DB()
	err = a.Close()
	if err != nil {
		fmt.Println(err)
	}
}
