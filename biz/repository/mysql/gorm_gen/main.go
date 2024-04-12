package main

import (
	"OceanShipBeidouMS/biz/config"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		//  设置输出路径
		OutPath: "./biz/repository/mysql/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, // 选择生成模式
	})
	//  建立数据库连接
	db := conn()
	g.UseDB(db) // 选择数据库连接

	g.ApplyBasic(
		// 从当前数据库的所有表生成结构
		g.GenerateAllTable()...,
	)
	// 生成代码
	g.Execute()
}

func conn() *gorm.DB {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetConfig().Mysql.User, config.GetConfig().Mysql.Password, config.GetConfig().Mysql.TCP, config.GetConfig().Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		hlog.Fatalf("connect mysql error:[%#v]\n", err)
		return nil
	}
	return db
}
