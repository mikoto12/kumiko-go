package consumer

import (
	"kumiko/pkg/logger"
	"kumiko/pkg/rabbitmq"
)

func TestQueue() {
	msgs, err := rabbitmq.Consume("test_queue")
	if err != nil {
		logger.StdError("监听test_queue失败: %v", err)
		return
	}
	for msg := range msgs {
		logger.StdInfo("[test_queue] 收到消息: %s", string(msg.Body))
		// 业务处理...
	}
}
