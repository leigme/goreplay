package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strings"
)

// RocketMqOutput is used for sending payloads to kafka in JSON format.
type RocketMqOutput struct {
	config   *OutputRocketMqConfig
	producer rocketmq.Producer
}

// NewRocketMqOutput creates instance of kafka producer client  with TLS config
func NewRocketMqOutput(address string, config *OutputRocketMqConfig) PluginWriter {

	var err error

	var p rocketmq.Producer

	p, err = rocketmq.NewProducer(
		producer.WithNameServer(strings.Split(config.Host, ",")),
		producer.WithRetry(config.Retry),
		producer.WithGroupName(config.Group),
	)

	if err != nil {
		fmt.Printf("create producer error: %s", err.Error())
		os.Exit(1)
	}

	err = p.Start()

	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	o := &RocketMqOutput{
		config:   config,
		producer: p,
	}

	return o
}

// PluginWrite writes a message to this plugin
func (o *RocketMqOutput) PluginWrite(msg *Message) (n int, err error) {

	var message *primitive.Message

	meta := strings.Split(string(msg.Meta), " ")

	rmm := RocketMqMessage{
		RecordId: o.config.RecordId,
		NodeIp:   o.config.NodeIp,
		Uuid:     meta[1],
		Type:     string(msg.Type),
		Data:     msg.Data,
		CapTime:  meta[2],
	}

	jsonData, err := json.Marshal(rmm)

	if err != nil {
		fmt.Println(err)
	}

	message = &primitive.Message{
		Topic: o.config.Topic,
		Body:  jsonData,
	}

	res, err := o.producer.SendSync(context.Background(), message)

	if err != nil {
		fmt.Printf("send message to rocketmq is error: %s", err.Error())
		os.Exit(1)
	}

	return int(res.Status), nil
}
