package main

import (
	_ "kumiko/docs"
	// "kumiko/internal/consumer"
	"kumiko/internal/router"
	"kumiko/pkg/database"
	// "kumiko/pkg/elasticsearch"
	"kumiko/pkg/logger"
	// "kumiko/pkg/rabbitmq"
	"kumiko/pkg/redis"

	"github.com/spf13/viper"
)

// @title kumiko API
// @version 1.0
// @description Gin + GORM 项目模板
// @host localhost:8080
// @BasePath /

func main() {
	// 初始化日志模块
	logger.InitLogger(viper.GetString("logger.level"), "")
	// 初始化数据库连接
	database.InitDB()
	// 初始化redis
	redis.InitRedis()
	// 初始化rabbitmq
	// rabbitmq.InitRabbitMQ()
	// 初始化es
	// elasticsearch.InitElasticsearch()
	// goroutine自动注册internal/consumer/下的所有消费者
	// for _, fn := range consumer.AllConsumers() {
	// 	go fn()
	// }
	r := router.SetupRouter()
	r.Run(":8080")
}
