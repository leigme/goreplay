package main

// OutputRocketMqConfig is the representation of rocketmq output configuration
type OutputRocketMqConfig struct {
	RecordId string `json:"output-rocketmq-record-id"`
	NodeIp   string `json:"output-rocketmq-node-ip"`
	Host     string `json:"output-rocketmq-host"`
	Retry    int    `json:"output-rocketmq-retry"`
	Group    string `json:"output-rocketmq-group"`
	Topic    string `json:"output-rocketmq-topic"`
}

// RocketMqMessage should contains catched request information that should be
// passed as Json to Apache RocketMq.
type RocketMqMessage struct {
	RecordId string `json:"record_id"`
	NodeIp   string `json:"node_ip"`
	Uuid     string `json:"uuid"`
	Type     string `json:"type"`
	Data     []byte `json:"data"`
	CapTime  string `json:"cap_time"`
}
