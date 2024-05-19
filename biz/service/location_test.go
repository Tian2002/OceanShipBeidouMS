package service

import (
	"math"
	"testing"
)

func TestLocation_haversineDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "夏威夷岛-所罗门群岛",
			args:    args{20.6886, -156.2358, -9.6500, 160.1833},
			want:    5836801,
			wantErr: false,
		},
	}
	// floatEqual 检查两个浮点数是否大致相等
	floatEqual := func(a, b, epsilon float64) bool {
		return math.Abs(a-b) < epsilon
	}
	//开始测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lo := &Location{}
			got, err := lo.haversineDistance(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2)
			if (err != nil) != tt.wantErr {
				t.Errorf("haversineDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !floatEqual(got, tt.want, 1) {
				t.Errorf("haversineDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
