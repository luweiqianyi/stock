package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"stock/cmd/transaction/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
