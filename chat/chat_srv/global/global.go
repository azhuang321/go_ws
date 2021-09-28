package global

import (
	"chat_srv/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	Config = &config.Config{}
)
