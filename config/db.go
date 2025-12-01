package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	IsConfigErr bool
	ConfigErr   error
)

func DBconnect() {
	var err error
	dsn := "host=localhost user=postgres password=466450 dbname=unittest port=5432"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}