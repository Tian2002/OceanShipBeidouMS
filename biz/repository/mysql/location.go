package mysql

import (
	"OceanShipBeidouMS/biz/repository/mysql/model"
	"context"
)

type LocationRepo interface {
	GetLocationByShipID(shipID int64) (*model.Location, error)
}

func GetLocationByShipID(ctx context.Context, shipID int64) (*model.Location, error) {
	return nil, nil
}
