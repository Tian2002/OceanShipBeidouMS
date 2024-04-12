package handler

import (
	"OceanShipBeidouMS/biz/common"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spf13/cast"
)

func ReceiveMessage(ctx context.Context, c *app.RequestContext) {
	stationParam := c.Param("stationMark")
	station := cast.ToInt(stationParam)
	if station <= 0 {
		c.JSON(returnError(common.ParamError))
	}

	err := c.BindForm(BdsData{})
	if err != nil {
		c.JSON(returnError(common.ParamError))
	}

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

// IReceiveMessage 接收消息处理
type IReceiveMessage interface {
	IBeiDouCard
	IProcessingLocationInformation
	// ProcessingMessage 解析北斗消息
	ProcessingMessage(originatorCardID int, destinationCardID int, timestamp int64, messageContent string) error
}

// IBeiDouCard 北斗卡号处理
type IBeiDouCard interface {
	// CheckStation 检测来源的基站与北斗卡号对应基站是否一致
	CheckStation(originatorID string, station int) error
	// CheckMessageType 检查目标卡号是组还是个人，并于传来的MessageType校对
	CheckMessageType(cardId string, messageType string) error
	// ParseBeiDouCard 通过卡号得到对应终端的携带者
	ParseBeiDouCard(cardID string) (shipID int, err error)
}

// IProcessingLocationInformation 位置信息处理
type IProcessingLocationInformation interface {
	// ProcessingLocationInformation 处理位置信息，包括检查位置信息的合理性
	ProcessingLocationInformation(cardID string, longitude float64, latitude float64, timestamp int64) error
}
