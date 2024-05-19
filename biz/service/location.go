package service

import (
	"fmt"
	"math"
)

type Location struct {
}

func (l *Location) ProcessingLocationInformation(shipID int, longitude float64, latitude float64, timestamp int64) error {
	//TODO implement me
	return nil
}

// haversineDistance 计算两个地理坐标点之间的距离（单位：米）
func (*Location) haversineDistance(lat1, lon1, lat2, lon2 float64) (float64, error) {
	const earthRadiusKm = 6371 // 地球半径，单位：千米

	// 检查经纬度是否在有效范围内
	isValidLatitude := func(lat float64) bool {
		return lat >= -90 && lat <= 90
	}

	isValidLongitude := func(lon float64) bool {
		return lon >= -180 && lon <= 180
	}

	if !isValidLatitude(lat1) || !isValidLatitude(lat2) {
		return 0, fmt.Errorf("invalid latitude: must be between -90 and 90 degrees")
	}
	if !isValidLongitude(lon1) || !isValidLongitude(lon2) {
		return 0, fmt.Errorf("invalid longitude: must be between -180 and 180 degrees")
	}

	// 将经纬度转换为弧度
	var (
		radians     = math.Pi / 180
		phi1        = lat1 * radians
		phi2        = lat2 * radians
		deltaPhi    = (lat2 - lat1) * radians
		deltaLambda = (lon2 - lon1) * radians
	)

	// 计算哈弗辛公式
	var (
		a = math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
			math.Cos(phi1)*math.Cos(phi2)*
				math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
		c        = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		distance = earthRadiusKm * c // 计算距离，单位：千米
	)

	// 将距离转换为米
	return distance * 1000, nil
}
