package global

import (
	"gorm.io/gorm"
	"user_api/config"
)

var (
	DB *gorm.DB

	Config = &config.Config{}
)
