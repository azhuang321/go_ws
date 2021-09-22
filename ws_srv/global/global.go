package global

import (
	"gorm.io/gorm"
	"ws_srv/config"
)

var (
	DB *gorm.DB

	Config = &config.Config{}
)
