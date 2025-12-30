package consumer

import (
	"kumiko/pkg/logger"
	"kumiko/pkg/rabbitmq"
)

func UserQueue() {
	msgs, err := rabbitmq.Consume("user_queue")
	if err != nil {
		logger.StdError("监听user_queue失败: %v", err)
		return
	}
	for msg := range msgs {
		logger.StdInfo("[user_queue] 收到消息: %s", string(msg.Body))
		// 业务处理...
	}
}
