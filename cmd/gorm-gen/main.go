package main

import (
	"OceanShipBeidouMS/internal/repository/mysql"
	"context"
	"gorm.io/gen"
)

// generate code
func main() {

	db := mysql.GetDb()
	cfg := gen.Config{
		OutPath:       "../../internal/repository/mysql/gen",
		Mode:          gen.WithQueryInterface,
		FieldNullable: true,
	}

	g := gen.NewGenerator(cfg)

	g.UseDB(db.WithContext(context.Background()))

	g.ApplyBasic(g.GenerateAllTable())

	g.Execute()
}
