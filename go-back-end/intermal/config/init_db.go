package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=exchange password=123 dbname=exchangedb port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
