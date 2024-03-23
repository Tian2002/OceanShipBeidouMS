package mysql

import (
	"OceanShipBeidouMS/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db = new(gorm.DB)

func connect() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%stcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetConfig().Mysql.User, config.GetConfig().Mysql.Password, config.GetConfig().Mysql.TCP, config.GetConfig().Mysql.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "bd_", // 定义表前缀，没有全局定义时，生成时需要输写完整表名
			SingularTable: true,  // 统一采用单数的时候定义结构名
		},
	})
	if err != nil {
		log.Fatalf("connect mysql error:[%#v]\n", err)
	}
}

func GetDb() *gorm.DB {
	return db
}
