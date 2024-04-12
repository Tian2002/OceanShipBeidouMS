package handler

import (
	"OceanShipBeidouMS/biz/common"
	"OceanShipBeidouMS/biz/service"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/cast"
	"strings"
)

func ReceiveMessage(ctx context.Context, c *app.RequestContext) {
	//router: /app/message2mis/station/:stationMark
	//解析stationMark参数
	stationParam := c.Param("stationMark")
	station := cast.ToInt(stationParam)
	if station <= 0 {
		c.JSON(returnError(common.ParamError))
		return
	}

	//解析post携带的参数
	bdsData := new(BdsData)
	if err := c.BindForm(bdsData); err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}
	if err := bdsData.Check(); err != nil {
		c.JSON(returnError(common.ParamError, err.Error()))
		return
	}

	//具体处理
	var receiveMessageService IReceiveMessage = &service.ReceiveMessage{
		BeiDouCard: new(service.BeiDouCard),
		Location:   new(service.Location),
	}

	//先做卡号相关操作
	if err := receiveMessageService.CheckStation(bdsData.OriginatorID, station); err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}
	if err := receiveMessageService.CheckMessageType(bdsData.DestinationID, bdsData.MessageType); err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}
	origin, err := receiveMessageService.ParseBeiDouCard(bdsData.OriginatorID)
	if err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}
	destination, err := receiveMessageService.ParseBeiDouCard(bdsData.DestinationID)
	if err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}

	//再做位置的操作
	if err := receiveMessageService.ProcessingLocationInformation(origin, bdsData.Longitude, bdsData.Latitude, bdsData.Timestamp); err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}

	//最后去处理消息
	if err := receiveMessageService.ProcessingMessage(origin, destination, bdsData.Timestamp, bdsData.MessageContent); err != nil {
		c.JSON(returnError(common.ParamError))
		return
	}

	c.JSON(returnSuccess(nil))
	return
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
	CheckMessageType(destinationID string, messageType string) error
	// ParseBeiDouCard 通过卡号得到对应终端的携带者
	ParseBeiDouCard(cardID string) (shipID int, err error)
}

// IProcessingLocationInformation 位置信息处理
type IProcessingLocationInformation interface {
	// ProcessingLocationInformation 处理位置信息，包括检查位置信息的合理性
	ProcessingLocationInformation(shipID int, longitude float64, latitude float64, timestamp int64) error
}

type ICheckReceiveMessageParam interface {
	Check() error
}

// Check 检查BdsData中每个字段合法性
func (bdsData *BdsData) Check() error {
	if bdsData.OriginatorID == "" {
		err := errors.New("checkBdsData error:OriginatorID is empty")
		hlog.Error(err)
		return err
	}
	if bdsData.DestinationID == "" {
		err := errors.New("checkBdsData error:DestinationID is empty")
		hlog.Error(err)
		return err
	}
	if bdsData.MessageType == "" {
		err := errors.New("checkBdsData error:MessageType is empty")
		hlog.Error(err)
		return err
	}
	bdsData.MessageType = strings.ToLower(bdsData.MessageType) //将message_type转换成小写,方便后续判定
	if bdsData.MessageType != common.MessageTypeGroup && bdsData.MessageType != common.MessageTypePerson {
		err := errors.New("checkBdsData error:message_type is not group or person")
		hlog.Error(err)
		return err
	}
	if bdsData.Timestamp == 0 {
		err := errors.New("checkBdsData error:timestamp is empty")
		hlog.Error(err)
		return err
	}
	if bdsData.MessageID == 0 {
		err := errors.New("checkBdsData error:message_id is empty")
		hlog.Error(err)
		return err
	}
	if bdsData.Latitude == 0 {
		err := errors.New("checkBdsData error:latitude is empty")
		hlog.Error(err)
		return err
	}
	if bdsData.Longitude == 0 {
		err := errors.New("checkBdsData error:longitude is empty")
		hlog.Error(err)
		return err
	}
	////允许发送空的消息，仅是日常消息的上报
	//if bdsData.MessageContent == "" {
	//}
	return nil
}
