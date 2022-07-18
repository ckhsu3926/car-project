package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(config MysqlConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
	mysqlClient, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		log.Fatalf("database connect failed: %v", err)
	}

	if mysqlClient.Error != nil {
		log.Fatalf("database error: %v", err)
	}

	return mysqlClient
}
