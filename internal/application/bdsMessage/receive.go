package bdsMessage

import (
	"OceanShipBeidouMS/internal/common"
	"OceanShipBeidouMS/internal/repository/kafkaRepo"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

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

// ParseBdsData 解析kafka发来的数据，通过一个channel持续返回解析后的数据
func ParseBdsData(ctx context.Context, topic string, groupId string, ch chan<- *BdsData) {
	reader := kafkaRepo.Topic(topic).NewKafkaReader(ctx, groupId)
	defer reader.Close()
	//一个无限循环读取数据发送到channel中进一步处理
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("ParseBdsData error:[%#v]", err)
			continue
		}
		bdsData := &BdsData{}
		err = bdsData.unmarshal(m.Value)
		if err != nil {
			log.Printf("ParseBdsData error:[%#v]", err)
			continue
		}
		err = bdsData.checkBdsData()
		if err != nil {
			log.Printf("ParseBdsData error:[%#v]", err)
			continue
		}
		ch <- bdsData
	}
}

// CheckBdsData 检查BdsData中每个字段合法性
func (bdsData *BdsData) checkBdsData() error {
	if bdsData.OriginatorID == "" {
		err := errors.New("originator_id is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.DestinationID == "" {
		err := errors.New("destination_id is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.MessageType == "" {
		err := errors.New("message_type is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	bdsData.MessageType = strings.ToLower(bdsData.MessageType) //将message_type转换成小写,方便后续判定
	if bdsData.MessageType != common.MessageTypeGroup && bdsData.MessageType != common.MessageTypePerson {
		err := errors.New("message_type is not group or person")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.Timestamp == 0 {
		err := errors.New("timestamp is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.MessageID == 0 {
		err := errors.New("message_id is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.Latitude == 0 {
		err := errors.New("latitude is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	if bdsData.Longitude == 0 {
		err := errors.New("longitude is empty")
		log.Printf("checkBdsData error:[%#v]", err)
		return err
	}
	////允许发送空的消息，仅是日常消息的上报
	//if bdsData.MessageContent == "" {
	//}
	return nil
}

// unmarshal 将kafka发来的数据解析成BdsData
func (bdsData *BdsData) unmarshal(data []byte) error {
	err := json.Unmarshal(data, bdsData)
	if err != nil {
		log.Printf("unmarshal error:[%#v]", err)
		return err
	}
	return nil
}
