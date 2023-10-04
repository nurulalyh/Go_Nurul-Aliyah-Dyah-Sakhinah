package database

import (
	"clean_arc/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c config.Config) *gorm.DB {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Cannot connect to database, ", err.Error())
	}

	return db
}
