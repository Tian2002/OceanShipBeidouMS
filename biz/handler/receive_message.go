package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func ReceiveMessage(ctx context.Context, c *app.RequestContext) {
	
}

type BdsData struct {
	//基本消息信息
	OriginatorID  string `json:"originator_id"`  // 发送方ID
	DestinationID string `json:"destination_id"` // 接收方ID,接收方id包含TalkGroupID，也就是发送的组消息，通过messageType来区分
	MessageType   string `json:"message_type"`   // 消息类型,标记是组消息还是个人消息
	Timestamp     int64  `json:"timestamp"`      // 时间戳
	MessageID     int64  `json:"message_id"`     // 消息ID
	//位置信息，经纬度
	Longitude float64 `json:"longitude"` // 经度
	Latitude  float64 `json:"latitude"`  // 纬度
	// 消息内容，下层是直接将消息使用string字符串发来，我们再对这个数据解析
	MessageContent string `json:"message_content"`
}
