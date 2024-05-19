package service

import (
	"OceanShipBeidouMS/biz/common"
	"OceanShipBeidouMS/biz/repository/mysql/model"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

type ReceiveMessage struct {
	*BeiDouCard
	*Location
}

func (r *ReceiveMessage) ProcessingMessage(originatorCardID int, destinationCardID int, timestamp int64, messageContent string) error {
	//TODO implement me
	return nil
}

func (r *ReceiveMessage) CheckLocationMessage(ctx context.Context, shipID int64, lat, lon float64, t int64) error {
	//now := time.Now().Unix()
	now := int64(1715702400)
	//location, err := mysql.GetLocationByShipID(ctx, shipID)
	//if err!=nil{
	//	return err
	//}
	location := model.Location{
		Longitude: 179,
		Latitude:  0,
		Timestamp: time.Unix(1715616000, 0),
	}
	lastTime := location.Timestamp.Unix()
	//检验时间戳
	if t <= lastTime || t > now {
		err := errors.New("timestamp illegal")
		hlog.Errorf("check location message error:[%v],now:[%v],lastTime:[%v],timestamp:[%v]", err, now, lastTime, t)
		return err
	}

	//经纬度数值的校验在haversineDistance函数中
	distance, err := r.haversineDistance(location.Latitude, location.Longitude, lat, lon)
	if err != nil {
		return err
	}

	// 计算时间差，单位为小时
	duration := float64((t - lastTime)) / 60 / 60
	// 将距离转换为海里，1海里约等于1852米
	distanceInNauticalMiles := float64(distance) / 1852.0
	// 计算速度，单位为海里/小时(节)
	speedInKnots := distanceInNauticalMiles / duration
	if speedInKnots > common.MaxShipSpeedInKnots {
		err := errors.New("location illegal")
		hlog.Errorf("check location message error:[%v],speedInKnots:[%v]", err, speedInKnots)
		return err
	}
	return nil
}
