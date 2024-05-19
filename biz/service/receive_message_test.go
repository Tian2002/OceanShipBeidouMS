package service

import (
	"context"
	"fmt"
	"testing"
)

func TestReceiveMessage_CheckLocationMessage(t *testing.T) {
	type fields struct {
		BeiDouCard *BeiDouCard
		Location   *Location
	}
	type args struct {
		ctx    context.Context
		shipID int64
		lat1   float64
		lon1   float64
		t      int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		//等价类划分的数据
		{
			args: args{
				lat1: 0,
				lon1: 179.001,
				t:    1715659200,
			},
			wantErr: false,
		},
		{
			args: args{
				lat1: 4,
				lon1: 181,
				t:    1715659230,
			},
		},
		{
			args: args{
				lat1: -91,
				lon1: 178.998,
				t:    1715659215,
			},
		},
		{
			args: args{
				lat1: 0,
				lon1: 179.001,
				t:    1715702401,
			},
		},
		{
			args: args{
				lat1: 3,
				lon1: 166,
				t:    1715659206,
			},
		},
		//边值测试数据
		{
			args: args{
				lat1: 0.01,
				lon1: 179.001,
				t:    1715616000,
			},
		},
		{
			args: args{
				lat1: 0.01,
				lon1: 179.001,
				t:    1715702400,
			},
		},
	}

	//开始测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReceiveMessage{}
			if err := r.CheckLocationMessage(tt.args.ctx, tt.args.shipID, tt.args.lat1, tt.args.lon1, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("data:[%v],CheckLocationMessage() error = %v", fmt.Sprintf("经度：[%v],纬度：[%v],时间戳:[%v]", tt.args.lon1, tt.args.lat1, tt.args.t), err)
				return
			}
			t.Logf("data:[%v],CheckLocationMessage() success", fmt.Sprintf("经度：[%v],纬度：[%v],时间戳:[%v]", tt.args.lon1, tt.args.lat1, tt.args.t))
		})
	}
}
