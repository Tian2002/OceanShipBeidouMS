package kafkaRepo

import (
	"OceanShipBeidouMS/internal/config"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type kafkaService interface {
	WriteKafka(ctx context.Context, key []byte, value []byte) (err error)
	ReadKafka(ctx context.Context, groupId string) (Key []byte, Value []byte, err error)
	NewKafkaReader(ctx context.Context, groupId string) (reader *kafka.Reader)
	SetTimeout(time time.Duration) *kafkaCli
	SetRetryCount(retryCount int) *kafkaCli
}

type kafkaCli struct {
	addr           []string
	topic          string
	writeTimeout   time.Duration
	readTimeout    time.Duration
	retryCount     int
	commitInterval time.Duration
	startOffset    int64
}

func defaultKafkaCli() *kafkaCli {
	return &kafkaCli{
		addr:           config.GetConfig().Kafka.Addr,
		writeTimeout:   time.Duration(config.GetConfig().Kafka.WriteTimeout) * time.Second,
		readTimeout:    time.Duration(config.GetConfig().Kafka.ReadTimeout) * time.Second,
		retryCount:     config.GetConfig().Kafka.RetryCount,
		commitInterval: time.Duration(config.GetConfig().Kafka.CommitInterval) * time.Second,
		startOffset:    int64(config.GetConfig().Kafka.StartOffset),
	}
}

// Topic 选择操作那个topic，在通过后续的函数调用选择具体操作
func Topic(topic string) *kafkaCli {
	cli := defaultKafkaCli()
	cli.topic = topic
	return cli
}

// WriteKafka 向kafka写入数据
func (cli *kafkaCli) WriteKafka(ctx context.Context, key []byte, value []byte) (err error) {
	writer := kafka.Writer{
		Addr:                   kafka.TCP(cli.addr...),
		Topic:                  cli.topic,
		WriteTimeout:           cli.writeTimeout,
		Balancer:               &kafka.Hash{},
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}
	defer writer.Close()

	for i := 0; i < cli.retryCount; i++ {
		err = writer.WriteMessages(ctx, kafka.Message{Key: key, Value: value})
		if err == nil {
			log.Printf("writeKafka success,topic:[%#v],key:[%#v],value:[%#v]\n", cli.topic, key, value)
			return nil
		}
	}
	if err != nil {
		log.Printf("writeKafka Error:[%#v]", err)
		return err
	}
	return
}

// ReadKafka 从kafka读取一条数据
func (cli *kafkaCli) ReadKafka(ctx context.Context, groupId string) (Key []byte, Value []byte, err error) {
	reader := kafka.NewReader(kafka.ReaderConfig{Brokers: cli.addr, Topic: cli.topic, CommitInterval: cli.commitInterval, GroupID: groupId, StartOffset: cli.startOffset})
	defer reader.Close()

	for i := 0; i < config.GetConfig().Kafka.RetryCount; i++ {
		message, err := reader.ReadMessage(ctx)
		if err == nil {
			log.Printf("readKafka success,message:[%#v]", message)
			return message.Key, message.Value, err
		}
	}

	if err != nil {
		log.Printf("readKafka error:[%#v]", err)
	}
	return
}

// NewKafkaReader 从kafka读取数据，返回一个reader，需要调用者使用完成后自己关闭这个reader
func (cli *kafkaCli) NewKafkaReader(ctx context.Context, groupId string) (reader *kafka.Reader) {
	return kafka.NewReader(kafka.ReaderConfig{Brokers: cli.addr, Topic: cli.topic, CommitInterval: cli.commitInterval, GroupID: groupId, StartOffset: cli.startOffset})
}

// SetTimeout 修改这次操作的默认writeTimeout、readTimeout配置字段
func (cli *kafkaCli) SetTimeout(time time.Duration) *kafkaCli {
	cli.writeTimeout = time
	cli.readTimeout = time

	return cli
}

// SetRetryCount 修改这次操作的默认retryCount配置字段
func (cli *kafkaCli) SetRetryCount(retryCount int) *kafkaCli {
	cli.retryCount = retryCount
	return cli
}

// SetCommitInterval 修改这次操作的默认commitInterval时间
func (cli *kafkaCli) SetCommitInterval(commitInterval time.Duration) *kafkaCli {
	cli.commitInterval = commitInterval
	return cli
}

// SetStartOffset 修改这次操作的默认startOffset
func (cli *kafkaCli) SetStartOffset(startOffset int64) *kafkaCli {
	cli.startOffset = startOffset
	return cli
}
