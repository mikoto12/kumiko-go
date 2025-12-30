package database

import (
	"kumiko/pkg/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.StdError("Error reading config file: %v", err)
	}
	dsn := viper.GetString("database.dsn")
	logger.StdInfo("Using DSN:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.StdError("failed to connect database")
	}
}
